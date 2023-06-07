package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"strings"
)

// Заголовок (Item.Title) - ключ
func HandleDel(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != models.LocalIp {
		sendResponse(w, models.CreateErrResponse(1))
		return
	}

	body, err := parseRequestBody(r)

	if err != nil {
		sendResponse(w, models.CreateErrResponse(2))
		return
	}

	if body.Title == "" {
		sendResponse(w, models.CreateErrResponse(4))
	} else {
		data := models.GetData()
		for i := 0; i < len(data.Items); i++ {
			if data.Items[i].Title == body.Title {
				data.Items = append(data.Items[:i], data.Items[i+1:]...)
				i--
			}
		}
		data.Write()
		sendResponse(w, models.CreateOkResponse())
	}
}
