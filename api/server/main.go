package main

import (
	"log"

	"github.com/ady243/go-project-chat/db"
	"github.com/ady243/go-project-chat/internal/user"
	websocket "github.com/ady243/go-project-chat/internal/webSocket"
	"github.com/ady243/go-project-chat/router"
)

func main() {

	db, err := db.NewDB()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	userRepo := user.NewRepository(db.GetDB())
	userSrvc := user.NewService(userRepo)
	userHandler := user.Newhandler(userSrvc)

	hub := websocket.NewHub()
	wsHandler := websocket.NewHandler(hub)
	go hub.Run()

	router.InitRoutes(userHandler, wsHandler)
	router.Start(":8080")

}
