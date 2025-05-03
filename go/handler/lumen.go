package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/govision-hackvidia/hackvidia-backend/service"
)

type LumenResponse struct {
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
}

func (h *Handler) LumenHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		grpc_client := service.NewMllmClient()
		input_text := r.FormValue("text")
		if input_text == "" {
			w.WriteHeader(http.StatusBadRequest)
		}
		text := grpc_client.Chat(input_text, nil)
		w.Header().Set("Content-Type", "application/json")

		data, err := json.Marshal(LumenResponse{
			Text: text,
		})
		if err != nil {
			log.Println("unable to marshal JSON: ", err)
		}
		w.Write(data)
	}
}
