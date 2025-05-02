package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", r)
}
