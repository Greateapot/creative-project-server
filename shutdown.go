package main

import (
	"net/http"
	"strings"
	"time"
)

func HandleShutdown(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != local_ip {
		responseString(w, http.StatusForbidden, "access denied")
		return
	}

	responseString(w, http.StatusOK, "ok")

	go func() {
		time.Sleep(time.Second) // 1 sec delay for request
		server.Shutdown(r.Context())
	}()
}
