package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/govision-hackvidia/hackvidia-backend/handler"
)

func main() {
	host := os.Getenv("SERVICE_HOST")
	if host == "" {
		log.Fatal("SERVICE_HOST not set")
	}
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		log.Fatal("SERVICE_PORT not set")
	}
	r := mux.NewRouter()
	h := handler.NewHandler()
	r.HandleFunc("/api/v1/login", h.LoginHandler()).Methods("POST")
	r.HandleFunc("/api/v1/enviroscan", h.EnviroscanHandler()).Methods("POST")
	r.HandleFunc("/api/v1/lumen", h.LumenHandler()).Methods("POST")

	log.Printf("Listening on %s:%s", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r)
}
