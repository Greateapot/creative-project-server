package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"strconv"
	"strings"
)

func HandleConfig(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != local_ip {
		sendResponse(w, models.CreateErrResponse(0x01, "access denied"))
		return
	}

	changes := false
	config := models.GetConfig()

	DataFileName := r.FormValue("dataFileName")
	if DataFileName != "" {
		config.DataFileName = DecodeB64(DataFileName)
		changes = true
	}

	Port := r.FormValue("port")
	if Port != "" {
		// TODO: CP scanDelay
		config.Port = Port
		changes = true
	}

	ScanDelay := r.FormValue("scanDelay")
	if ScanDelay != "" {
		if v, err := strconv.Atoi(ScanDelay); err != nil {
			sendResponse(w, models.CreateErrResponse(0xC1, "err scanDelay"))
		} else {
			// TODO: assert max int val
			config.ScanDelay = v
			changes = true
		}
	}

	if !changes {
		sendResponse(w, models.CreateErrResponse(0xC2, "no changes"))
	} else {
		config.Write()
		sendResponse(w, models.CreateOkResponse())
	}

}
