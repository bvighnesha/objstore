package main

import (
	"objstore/console"
	"objstore/db"
)

func main() {
	db := db.NewObjectDB()
	console.Start(db)
}
