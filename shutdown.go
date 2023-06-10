package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"time"
)

func HandleShutdown(s *http.Server, w http.ResponseWriter, r *http.Request) {
	if !IsHostRequest(r) {
		sendResponse(w, models.GetResponseErrAccessDenied())
	} else {
		sendResponse(w, models.GetResponseOK())

		go func() {
			time.Sleep(time.Second) // 1 sec delay for request
			s.Shutdown(r.Context())
		}()
	}

}
