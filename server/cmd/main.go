package main

import (
	"log"
	"server/db/loaders"
)

func main() {
	log.Println("Starting connecting to db")

	loaders.ConnectToDb()

	log.Println("Db connected successfully")

}
