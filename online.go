package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"strings"
)

func HandleOnline(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != models.LocalIp {
		sendResponse(w, models.CreateErrResponse(1))
	} else {
		sendResponse(w, models.CreateDataResponse(models.GetOnline()))
	}
}
