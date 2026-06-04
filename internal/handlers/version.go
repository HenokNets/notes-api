package handlers

import (
	"encoding/json"
	"net/http"
)

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"version": "0.1.0",
	}

	json.NewEncoder(w).Encode(response)
}
