package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	AuthToken string
}

func (h *Handler) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Login Implementation
		bodyByte, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal("failed to read body: ", err)
		}
		var login_request LoginRequest
		if err := json.Unmarshal(bodyByte, &login_request); err != nil {
			log.Println("failed to unmarshal: ", err)
		}

		data, err := json.Marshal(LoginResponse{
			AuthToken: "123",
		})
		if err != nil {
			log.Println("error on marshal JSON: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write(data)
	}
}
