package main

import (
	"encoding/json"
	"greateapot/creative_project_server/models"
	"net/http"
	"strings"
)

func sendResponse(w http.ResponseWriter, statusCode int, r *models.Response) {
	if data, err := json.Marshal(r); err == nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(data))
	} else {
		// Не должно вернуться
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func IsHostRequest(r *http.Request) bool {
	return strings.Split(r.RemoteAddr, ":")[0] == models.LocalIp
}
