package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	if IsHostRequest(r) {
		sendResponse(w, http.StatusOK, *models.GetItems())
	} else {
		sendResponse(w, http.StatusOK, *models.GetItems().HidePath())
	}
}
