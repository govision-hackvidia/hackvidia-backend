package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/govision-hackvidia/hackvidia-backend/service"
)

type EnviroscanResponse struct {
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
}

func (h *Handler) EnviroscanHandler() http.HandlerFunc {
	grpc_client := service.NewGrpcClient()
	return func(w http.ResponseWriter, r *http.Request) {
		grpc_client.Chat()
		// grpc_client.Client
		// flusher, _ := w.(http.Flusher)
		w.Header().Set("Content-Type", "application/json")
		// Simulate Chunk Stream

		text := "Hello, I am Enviroscan. Nice to meet you."
		data, err := json.Marshal(EnviroscanResponse{
			Text: text,
		})
		if err != nil {
			log.Println("unable to marshal JSON: ", err)
		}
		w.Write(data)
	}
}
