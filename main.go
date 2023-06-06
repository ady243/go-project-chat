package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	setupRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRoutes() {
	http.HandleFunc("/users", createUserHandler)
	http.HandleFunc("/users/{id}", deleteUserHandler)
	http.HandleFunc("/rooms", createRoomHandler)
	http.HandleFunc("/rooms/{id}", deleteRoomHandler)
	http.HandleFunc("/rooms/{id}/messages", sendMessageHandler)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "User created")
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "User deleted")
}

func createRoomHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Room created")
}

func deleteRoomHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Room deleted")
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Message sent")
}
