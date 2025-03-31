package main

import (
	"log"
	"server/db/loaders"
)

func main() {
	log.Println("Starting connecting to db")

	db, err := loaders.ConnectToDb()

	if err != nil {
		log.Fatalf("❌ Could not connect to database: %v", err)
	}
	
	defer db.Close()

	log.Println("Db connected successfully")

}
