package handler

import (
	"log"

	"github.com/govision-hackvidia/hackvidia-backend/db"
)

type Handler struct {
	DB *db.DB
}

func NewHandler() *Handler {
	db_conn := db.NewDB()
	if db_conn == nil {
		log.Fatal("unable to connect psql")
	}
	return &Handler{
		DB: db_conn,
	}
}
