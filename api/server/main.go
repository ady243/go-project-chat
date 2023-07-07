package main

import (
	"log"

	"github.com/ady243/go-project-chat/db"
)

func main() {
	_, err := db.NewDB()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

}
