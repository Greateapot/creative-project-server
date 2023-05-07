package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"strings"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] == local_ip {
		sendResponse(w, models.CreateDataResponse(*models.GetData()))
	} else {
		sendResponse(w, models.CreateDataResponse(*models.GetHiddenData()))
	}
}
