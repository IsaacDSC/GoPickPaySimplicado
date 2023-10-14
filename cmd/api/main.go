package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type TransactionInputDto struct {
	Value      float64 `json:"value"`
	SenderID   string  `json:"senderId"`
	ReceiverID string  `json:"receiverId"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("alive"))
	})

	http.HandleFunc("/api/v1/transaction", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		payload, _ := io.ReadAll(r.Body)
		var body TransactionInputDto
		err := json.Unmarshal(payload, &body)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(payload)
	})

	http.ListenAndServe(":3000", nil)
}
