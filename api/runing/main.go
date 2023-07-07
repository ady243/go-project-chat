package main

import (
	"api/db"
	"log"
)

func main() {
	_, err := db.NewDB()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

}
