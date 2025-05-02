package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/govision-hackvidia/hackvidia-backend/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/login", handler.LoginHandler()).Methods("POST")
	r.HandleFunc("/api/v1/enviroscan", handler.EnviroscanHandler()).Methods("POST")
	r.HandleFunc("/api/v1/lumen", handler.LumenHandler()).Methods("POST")

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", r)
}
