package main

import (
	"fmt"
	"objstore/config"
	"objstore/console"
	"objstore/db"
)

func main() {
	host, port := config.ReadConfig()
	db := db.NewObjectDB(fmt.Sprintf("%s:%s", host, port))
	console.Start(db)

	/*err := db.Store(context.Background(), &obj.Person{
		Name:      "test5",
		LastName:  "Bojja",
		Birthday:  "04-11-1989",
		BirthDate: time.Now(),
	})*/

	//fmt.Println(err)

	/*obj, err := db.ListObjects(context.Background(), "person")
	fmt.Println("ListObjects", err)
	fmt.Println(obj)
	obj1, err := db.GetObjectByName(context.Background(), "test")
	fmt.Println("GetObjectByName", err)
	fmt.Println(obj1.GetName())*/
	//err = db.DeleteObject(context.Background(), obj[0].GetID())
	//fmt.Println("DeleteObject", err)
}
