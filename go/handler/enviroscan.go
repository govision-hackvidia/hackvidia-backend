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
		input_text := r.FormValue("text")
		if input_text == "" {
			w.WriteHeader(http.StatusBadRequest)
		}
		text := grpc_client.Chat(input_text, nil)
		w.Header().Set("Content-Type", "application/json")
		// Simulate Chunk Stream

		data, err := json.Marshal(EnviroscanResponse{
			Text: text,
		})
		if err != nil {
			log.Println("unable to marshal JSON: ", err)
		}
		w.Write(data)
	}
}
