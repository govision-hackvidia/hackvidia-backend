package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/govision-hackvidia/hackvidia-backend/handler"
)

func main() {
	r := mux.NewRouter()
	h := handler.NewHandler()
	r.HandleFunc("/api/v1/login", h.LoginHandler()).Methods("POST")
	r.HandleFunc("/api/v1/enviroscan", h.EnviroscanHandler()).Methods("POST")
	r.HandleFunc("/api/v1/lumen", h.LumenHandler()).Methods("POST")

	log.Println("Listening on :8080")
	http.ListenAndServe("0.0.0.0:8080", r)
}
