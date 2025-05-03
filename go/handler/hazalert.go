package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/govision-hackvidia/hackvidia-backend/service"
)

type HazalertResponse struct {
	NearestDistance float32 `json:"nearest_distance"`
}

func (h *Handler) HazalertHandler() http.HandlerFunc {
	grpc_client := service.NewHazalertClient()
	return func(w http.ResponseWriter, r *http.Request) {
		imgForm, imgHeader, err := r.FormFile("img")
		if err != nil || imgHeader != nil {
			log.Println("error on form file: ", err)
		}
		imgByte, err := io.ReadAll(imgForm)
		if err != nil {
			log.Println("error on read file: ", err)
		}
		nearest_distance := grpc_client.Predict(imgByte)
		w.Header().Set("Content-Type", "application/json")

		data, err := json.Marshal(HazalertResponse{
			NearestDistance: nearest_distance,
		})
		if err != nil {
			log.Println("unable to marshal JSON: ", err)
		}
		w.Write(data)
	}
}
