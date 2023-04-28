package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"objstore/config"
	"objstore/console"
	"objstore/db"
)

func main() {
	host, port := config.ReadConfig()
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})
	db := db.NewObjectDB(client)
	console.Start(db)
}
