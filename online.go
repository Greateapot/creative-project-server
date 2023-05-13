package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"strings"
)

// TODO: [list for GetOnline]
func HandleOnline(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != models.LocalIp {
		sendResponse(w, models.CreateErrResponse(0x01, "access denied"))
	} else {
		sendResponse(w, models.CreateDataResponse(models.GetOnline()))
	}

}
