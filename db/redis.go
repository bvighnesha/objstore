package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"objstore/model"
)

type DB struct {
	client *redis.Client
}

// NewObjectDB returns a new ObjectDB.
func NewObjectDB(client *redis.Client) *DB {
	return &DB{
		client: client,
	}
}

// Store will store the object in the data store. The object will have a name
// and kind, and the Store method should create a unique ID.
func (db *DB) Store(ctx context.Context, object model.Object) error {
	if object.GetName() == "" {
		return fmt.Errorf("object name cannot be empty")
	}
	keys, err := db.client.Keys(ctx, "*:"+object.GetName()+":*").Result()
	if err != nil {
		return fmt.Errorf("failed to get object keys: %w", err)
	}
	if len(keys) > 0 {
		return fmt.Errorf("object already exists: %d", len(keys))
	}

	id := uuid.New().String()
	object.SetID(id)
	b, err := json.Marshal(object)
	if err != nil {
		return fmt.Errorf("failed to marshal object: %w", err)
	}
	key := fmt.Sprintf("%s:%s:%s", object.GetKind(), object.GetName(), id)
	fmt.Println(key)
	err = db.client.Set(ctx, key, b, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to store object in Redis: %w", err)
	}
	return nil
}

// GetObjectByID will retrieve the object with the provided ID.
// The object will be retrieved from Redis using the ID.
func (db *DB) GetObjectByID(ctx context.Context, id string) (model.Object, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}
	var object model.Object
	keys, err := db.client.Keys(ctx, "*:"+id).Result()
	if err != nil {
		return object, fmt.Errorf("failed to get object keys: %w", err)
	}
	if len(keys) != 1 {
		return object, fmt.Errorf("unexpected number of keys found: %d", len(keys))
	}
	b, err := db.client.Get(ctx, keys[0]).Bytes()
	if err != nil {
		return object, fmt.Errorf("failed to get object from Redis: %w", err)
	}

	var person model.Person
	err = json.Unmarshal(b, &person)
	if err != nil {
		var animal model.Animal
		err = json.Unmarshal(b, &animal)
		if err != nil {
			return object, fmt.Errorf("failed to unmarshal object: %w", err)
		}
		object = &animal

		return object, nil
	}
	object = &person

	return object, nil
}

// GetObjectByName will retrieve the object with the given name.
// The object will be retrieved from Redis using the name.
func (db *DB) GetObjectByName(ctx context.Context, name string) (model.Object, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	var object model.Object
	fmt.Println(name)
	keys, err := db.client.Keys(ctx, "*"+name+"*").Result()
	if err != nil {
		return object, fmt.Errorf("failed to get object keys: %w", err)
	}
	if len(keys) != 1 {
		return object, fmt.Errorf("unexpected number of keys found: %d", len(keys))
	}

	b, err := db.client.Get(ctx, keys[0]).Bytes()
	if err != nil {
		return object, fmt.Errorf("failed to get object from Redis: %w", err)
	}

	var ob map[string]interface{}
	err = json.Unmarshal(b, &ob)
	if err != nil {
		return object, fmt.Errorf("failed to unmarshal object: %w", err)
	}
	if ob["name"] == name {
		var person model.Person
		err = json.Unmarshal(b, &person)
		if err != nil {
			var animal model.Animal
			err = json.Unmarshal(b, &animal)
			if err != nil {
				return object, fmt.Errorf("failed to unmarshal object: %w", err)
			}
			object = &animal

			return object, nil
		}
		object = &person
		return object, nil
	}

	return object, nil
}

// ListObjects will return a list of all objects of the given kind.
// The objects will be retrieved from Redis using the kind.
func (db *DB) ListObjects(ctx context.Context, kind string) ([]model.Object, error) {
	if kind == "" {
		return nil, fmt.Errorf("kind cannot be empty")
	}
	var objects []model.Object
	keys, err := db.client.Keys(ctx, "*").Result()
	if err != nil {
		return objects, fmt.Errorf("failed to get object keys: %w", err)
	}

	for _, key := range keys {
		b, err := db.client.Get(ctx, key).Bytes()
		if err != nil {
			return objects, fmt.Errorf("failed to get object from Redis: %w", err)
		}

		switch kind {
		case "person":
			var person model.Person
			err = json.Unmarshal(b, &person)
			if err != nil {
				continue
			}

			objects = append(objects, &person)
		case "animal":
			var animal model.Animal
			err = json.Unmarshal(b, &animal)
			if err != nil {
				continue
			}
			objects = append(objects, &animal)

		}
	}
	return objects, nil
}

// DeleteObject will delete the object.
// The object will be deleted from Redis using the ID.
func (db *DB) DeleteObject(ctx context.Context, id string) error {
	obj, err := db.GetObjectByID(ctx, id)
	if err != nil {
		return fmt.Errorf("error object not found: %w", err)
	}
	err = db.client.Del(ctx, fmt.Sprintf("%s:%s:%s", obj.GetKind(), obj.GetName(), id)).Err()
	if err != nil {
		return fmt.Errorf("error deleting object: %w", err)
	}
	return nil
}
