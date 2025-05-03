package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/govision-hackvidia/hackvidia-backend/service"
)

type EnviroscanResponse struct {
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
}

func (h *Handler) EnviroscanHandler() http.HandlerFunc {
	grpc_client := service.NewMllmClient()
	return func(w http.ResponseWriter, r *http.Request) {
		input_text := r.FormValue("text")
		imgForm, fileHeader, err := r.FormFile("img")
		if err != nil || fileHeader != nil {
			log.Println("error on form file: ", err)
			return
		}
		imgByte, err := io.ReadAll(imgForm)
		if err != nil {
			log.Println("error on read file: ", err)
		}
		text := grpc_client.Chat(input_text, imgByte)
		w.Header().Set("Content-Type", "application/json")

		data, err := json.Marshal(EnviroscanResponse{
			Text: text,
		})
		if err != nil {
			log.Println("unable to marshal JSON: ", err)
		}
		w.Write(data)
	}
}
