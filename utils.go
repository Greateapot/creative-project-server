package main

import (
	"encoding/json"
	"greateapot/creative_project_server/models"
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, r *models.Response) {
	if data, err := json.MarshalIndent(r, "", "  "); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(data))
	} else {
		// никогда не вернется...
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func parseRequestBody(r *http.Request) (body *models.Request, err error) {
	body = models.NewRequest()
	bytes, _ := io.ReadAll(r.Body)
	err = json.Unmarshal(bytes, body)
	return
}
