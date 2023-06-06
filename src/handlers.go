package src

import (
	"fmt"
	"net/http"
)

func Homehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to my chat api!")
}
