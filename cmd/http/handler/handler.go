package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// JSONResponse default Response struct for JSON
type JSONResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// JSONWriter writer helper for writing JSON
// if got encode error, returns 500
func JSONWriter(w http.ResponseWriter, response *JSONResponse) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

// JSONReader reader helper for reading JSON
// if got decode error, returns 500
func JSONReader(w http.ResponseWriter, r io.Reader, v interface{}) {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}
