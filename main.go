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
}
