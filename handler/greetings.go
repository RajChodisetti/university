package handler

import (
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to our University")
}
