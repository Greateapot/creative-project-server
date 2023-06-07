package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"strings"
)

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != models.LocalIp {
		sendResponse(w, models.CreateErrResponse(1))
		return
	}

	item := &models.Item{}

	if body, err := parseRequestBody(r); err == nil {
		item.Path = body.Path
		item.Title = body.Title
		item.Type = body.Type
	} else {
		sendResponse(w, models.CreateErrResponse(2))
		return
	}

	if item.Path == "" {
		sendResponse(w, models.CreateErrResponse(3))
		return
	}

	if item.Title == "" {
		sendResponse(w, models.CreateErrResponse(4))
		return
	}

	if item.Type < 0 {
		sendResponse(w, models.CreateErrResponse(5))
		return
	}

	data := models.GetData()

	flag := false
	for i := 0; !flag && i < len(data.Items); i++ {
		flag = data.Items[i].Title == item.Title
	}

	if flag {
		sendResponse(w, models.CreateErrResponse(9))
	} else {
		data.Items = append(data.Items, *item)
		data.Write()
		sendResponse(w, models.CreateOkResponse())
	}
}
