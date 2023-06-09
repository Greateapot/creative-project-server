package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
)

func HandleOnline(w http.ResponseWriter, r *http.Request) {
	if !IsHostRequest(r) {
		sendResponse(w, models.GetResponseErrAccessDenied())
	} else {
		sendResponse(w, models.CreateDataResponse(models.GetOnline()))
	}
}
