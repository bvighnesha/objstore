package main

import (
	"context"
	"fmt"
	db2 "objstore/db"
	obj2 "objstore/obj"
	"time"
)

func main() {
	fmt.Println("Hello World")
	db := db2.NewObjectDB()
	err := db.Store(context.Background(), &obj2.Person{
		Name:      "test",
		LastName:  "Bojja",
		Birthday:  "04-11-1989",
		BirthDate: time.Time{},
	})

	fmt.Println(err)

	obj, err := db.ListObjects(context.Background(), "person")
	fmt.Println("ListObjects", err)
	fmt.Println(obj[0].GetName())
	obj1, err := db.GetObjectByName(context.Background(), "test")
	fmt.Println("GetObjectByName", err)
	fmt.Println(obj1.GetName())
	err = db.DeleteObject(context.Background(), obj[0].GetID())
	fmt.Println("DeleteObject", err)
}
