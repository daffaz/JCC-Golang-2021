package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseToJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-type", "application/json")
	dataInByte, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "Error occured...", http.StatusBadRequest)
	}

	w.WriteHeader(status)
	w.Write(dataInByte)
}
