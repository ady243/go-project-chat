package main

import (
	"log"

	"github.com/ady243/go-project-chat/db"
	"github.com/ady243/go-project-chat/router"
	"github.com/ady243/go-project-chat/user"
)

func main() {

	db, err := db.NewDB()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	userRepo := user.NewRepository(db.GetDB())
	userSrvc := user.NewService(userRepo)
	userHandler := user.Newhandler(userSrvc)

	router.InitRoutes(userHandler)
	router.Start(":8080")

}
