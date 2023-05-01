package main

import (
	"greateapot/creative-project-server/models"
	"net/http"
	"strings"
)

func HandleConfig(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != local_ip {
		responseString(w, http.StatusForbidden, "access denied")
		return
	}

	changes := false

	DataFileName := r.FormValue("dataFileName")
	if DataFileName != "" {
		models.GetConfig().DataFileName = DataFileName
		changes = true
	}

	Port := r.FormValue("port")
	if Port != "" {
		models.GetConfig().Port = Port
		changes = true
	}

	if changes {
		models.GetConfig().Write()
		responseString(w, http.StatusOK, "ok")
	} else {
		responseString(w, http.StatusBadRequest, "no changes")
	}

}
