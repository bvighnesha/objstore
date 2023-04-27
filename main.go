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

	/*err := db.Store(context.Background(), &model.Person{
		Name:      "test5",
		LastName:  "Bojja",
		Birthday:  "04-11-1989",
		BirthDate: time.Now(),
	})*/

	//fmt.Println(err)

	/*model, err := db.ListObjects(context.Background(), "person")
	fmt.Println("ListObjects", err)
	fmt.Println(model)
	obj1, err := db.GetObjectByName(context.Background(), "test")
	fmt.Println("GetObjectByName", err)
	fmt.Println(obj1.GetName())*/
	//err = db.DeleteObject(context.Background(), model[0].GetID())
	//fmt.Println("DeleteObject", err)
}
