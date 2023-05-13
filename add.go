package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"strconv"
	"strings"
)

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != models.LocalIp {
		sendResponse(w, models.CreateErrResponse(0x01, "access denied"))
		return
	}

	item := &models.Item{}

	item.Path = r.FormValue("path")
	if item.Path == "" {
		sendResponse(w, models.CreateErrResponse(0xA1, "no path"))
		return
	}

	if r.FormValue("type") == "" {
		sendResponse(w, models.CreateErrResponse(0xA2, "no type"))
		return
	}

	if t, err := strconv.ParseUint(r.FormValue("type"), 10, 8); err == nil {
		item.Type = uint8(t)
	} else {
		sendResponse(w, models.CreateErrResponse(0xA3, "err type"))
		return
	}

	item.Title = r.FormValue("title")
	if item.Title == "" {
		sendResponse(w, models.CreateErrResponse(0xA4, "no title"))
		return
	}

	data := models.GetData()

	flag := false
	for i := 0; !flag && i < len(data.Items); i++ {
		flag = data.Items[i].Title == item.Title
	}

	if flag {
		sendResponse(w, models.CreateErrResponse(0xA5, "title already exists"))
	} else {
		data.Items = append(data.Items, *item)
		data.Write()
		sendResponse(w, models.CreateOkResponse())
	}
}
