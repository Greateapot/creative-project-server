package main

import (
	"greateapot/creative-project-server/models"
	"net/http"
	"strings"
)

// Заголовок (Item.Title) - ключ
func HandleDel(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != local_ip {
		responseString(w, http.StatusForbidden, "access denied")
		return
	}

	if title := r.FormValue("title"); title == "" {
		responseString(w, http.StatusBadRequest, "no title")
	} else {
		data := models.GetData()
		for i := 0; i < len(data.Items); i++ {
			if data.Items[i].Title == title {
				data.Items = append(data.Items[:i], data.Items[i+1:]...)
				i--
			}
		}
		data.Write()
		responseString(w, http.StatusOK, "ok")
	}
}
