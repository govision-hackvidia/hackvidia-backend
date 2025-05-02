package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type LumenResponse struct {
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
}

func (h *Handler) LumenHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: EnviroScan Implementation
		flusher, _ := w.(http.Flusher)
		// Simulate Chunk Stream

		text := "Hello, I am LUMEN"
		textSplit := strings.Split(text, " ")
		for i, t := range textSplit {
			is_done := false
			if i+1 == len(textSplit) {
				is_done = true
			}
			data, err := json.Marshal(LumenResponse{
				Text:   t,
				IsDone: is_done,
			})

			if err != nil {
				log.Println("error on marshal JSON: ", err)
				return
			}
			w.Write(data)
			flusher.Flush()
			time.Sleep(10 * time.Millisecond)
		}

	}
}
