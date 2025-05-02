package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type LumenResponse struct {
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
}

func (h *Handler) LumenHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: EnviroScan Implementation
		data, err := json.Marshal(LumenResponse{
			Text:   "lumen_test",
			IsDone: false,
		})
		if err != nil {
			log.Println("error on marshal JSON: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write(data)
	}
}
