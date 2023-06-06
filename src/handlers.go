package handlers

import (
	"fmt"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "User created")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "User deleted")
}
