package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	if IsHostRequest(r) {
		sendResponse(w, models.CreateDataResponse(*models.GetData()))
	} else {
		sendResponse(w, models.CreateDataResponse(*models.GetData().HidePath()))
	}
}
