package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"time"
)

func HandleShutdown(s *http.Server, w http.ResponseWriter, r *http.Request) {
	if !IsHostRequest(r) {
		sendResponse(w, http.StatusForbidden, models.ErrAccessDenied())
	} else {
		sendResponse(w, http.StatusOK, nil)

		go func() {
			time.Sleep(time.Second) // 1 sec delay for request
			s.Shutdown(r.Context())
		}()
	}

}
