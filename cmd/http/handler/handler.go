package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type JSONResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error,omitempty"`
}

func JSONWriter(w http.ResponseWriter, response *JSONResponse) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

func JSONReader(w http.ResponseWriter, r io.Reader, v interface{}) {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}
