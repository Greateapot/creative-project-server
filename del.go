package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"strings"
)

// Заголовок (Item.Title) - ключ
func HandleDel(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != models.LocalIp {
		sendResponse(w, models.CreateErrResponse(0x01, "access denied"))
		return
	}

	if title := r.FormValue("title"); title == "" {
		sendResponse(w, models.CreateErrResponse(0xD1, "no title"))
	} else {
		data := models.GetData()
		for i := 0; i < len(data.Items); i++ {
			if data.Items[i].Title == title {
				data.Items = append(data.Items[:i], data.Items[i+1:]...)
				i--
			}
		}
		data.Write()
		sendResponse(w, models.CreateOkResponse())
	}
}
