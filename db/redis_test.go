package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"objstore/model"
	"testing"
	"time"
)
import "github.com/alicebob/miniredis/v2"

func TestStore(t *testing.T) {
	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	db := NewObjectDB(client)
	db.Store(context.Background(), &model.Person{
		Name:      "test",
		LastName:  "test",
		Birthday:  "123",
		BirthDate: time.Time{},
	})
	data, err := db.GetObjectByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, "test", data.GetName())
}

func TestGetObjectByID(t *testing.T) {
	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	db := NewObjectDB(client)
	db.Store(context.Background(), &model.Person{
		Name:      "test",
		LastName:  "test",
		Birthday:  "123",
		BirthDate: time.Time{},
	})
	data, err := db.GetObjectByName(context.Background(), "test")

	data, err = db.GetObjectByID(context.Background(), data.GetID())
	assert.Nil(t, err)
	assert.Equal(t, "test", data.GetName())
}

func TestGetObjectByName(t *testing.T) {
	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	db := NewObjectDB(client)
	db.Store(context.Background(), &model.Person{
		Name:      "test",
		LastName:  "test",
		Birthday:  "123",
		BirthDate: time.Time{},
	})
	data, err := db.GetObjectByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, "test", data.GetName())
}

func TestDeleteObject(t *testing.T) {
	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	db := NewObjectDB(client)
	db.Store(context.Background(), &model.Person{
		Name:      "test",
		LastName:  "test",
		Birthday:  "123",
		BirthDate: time.Time{},
	})
	data, err := db.GetObjectByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, "test", data.GetName())
	err = db.DeleteObject(context.Background(), data.GetID())
	assert.Nil(t, err)
	data, err = db.GetObjectByName(context.Background(), "test")
	assert.NotNil(t, err)
	assert.Empty(t, data)
}
