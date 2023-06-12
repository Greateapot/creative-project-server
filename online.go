package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
)

func HandleOnline(w http.ResponseWriter, r *http.Request) {
	if !IsHostRequest(r) {
		sendResponse(w, http.StatusForbidden, models.ErrAccessDenied())
	} else {
		sendResponse(w, http.StatusOK, models.GetOnline())
	}
}
