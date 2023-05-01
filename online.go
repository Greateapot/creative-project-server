package main

import (
	"net/http"
	"strings"
)

func HandleOnline(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != local_ip {
		responseString(w, http.StatusForbidden, "access denied")
	} else {
		online := GetOnline()
		responseString(w, http.StatusOK, online)
	}

}
