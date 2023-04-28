package console

import (
	"context"
	"encoding/json"
	"fmt"
	"objstore/db"
	"objstore/model"
	"strings"
)

// Start will start the console.
// The console will prompt the user for a command, and then prompt for the
// required data for that command.
// The console will then call the appropriate function on the DB.
func Start(db *db.DB) {
	fmt.Println("Please provide a command: SET, GET, LIST or DEL")
	var command string
	fmt.Scanln(&command)

	switch command {
	case "SET":
		fmt.Println("Please provide a object to set")
		var data string
		fmt.Scanln(&data)

		var object model.Object

		var person model.Person
		err := json.Unmarshal([]byte(data), &person)
		if err != nil {
			var animal model.Animal
			err := json.Unmarshal([]byte(data), &animal)
			if err != nil {
				fmt.Errorf("failed to unmarshal object: %w", err)
			}
			object = &animal

		}
		object = &person

		err = db.Store(context.Background(), object)
		if err != nil {
			fmt.Println("Error storing object:", err)
		}
	case "GET":
		fmt.Println("Do you want to get by name or ID?")
		var kind string
		fmt.Scanln(&kind)

		switch strings.ToUpper(kind) {
		case "NAME":
			fmt.Println("Please provide a name to get")
			var name string
			fmt.Scanln(&name)
			o, err := db.GetObjectByName(context.Background(), name)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(o)
		case "ID":
			fmt.Println("Please provide an ID to get")
			var id string
			fmt.Scanln(&id)
			o, err := db.GetObjectByID(context.Background(), id)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(o)

		default:
			fmt.Println("Unknown command:", kind)
		}

	case "LIST":
		fmt.Println("Please provide a kind to list")
		var kind string
		fmt.Scanln(&kind)
		objs, err := db.ListObjects(context.Background(), kind)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, obj := range objs {
			fmt.Println(obj)
		}

	case "DEL":
		fmt.Println("Please provide a id to delete")
		var id string
		fmt.Scanln(&id)
		err := db.DeleteObject(context.Background(), id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Deleted ID %s\n", id)

	default:
		fmt.Println("Unknown command:", command)
	}
}
