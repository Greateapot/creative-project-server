package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"strings"
	"time"
)

func HandleShutdown(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != local_ip {
		sendResponse(w, models.CreateErrResponse(0x01, "access denied"))
		return
	}

	sendResponse(w, models.CreateOkResponse())

	go func() {
		time.Sleep(time.Second) // 1 sec delay for request
		server.Shutdown(r.Context())
	}()
}
