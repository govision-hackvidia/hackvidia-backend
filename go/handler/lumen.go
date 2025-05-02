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
		// flusher, _ := w.(http.Flusher)
		// Simulate Chunk Stream

		text := "Hello, I am LUMEN"

		data, err := json.Marshal(LumenResponse{
			Text: text,
		})
		if err != nil {
			log.Println("unable to marshal JSON: ", err)
		}
		w.Write(data)

		// textSplit := strings.Split(text, " ")
		// for _, t := range textSplit {
		// w.Write([]byte(t))
		// flusher.Flush()
		// time.Sleep(10 * time.Millisecond)
		// }

	}
}
