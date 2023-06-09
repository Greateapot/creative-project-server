package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
)

func HandleDel(w http.ResponseWriter, r *http.Request) {
	if !IsHostRequest(r) {
		sendResponse(w, models.GetResponseErrAccessDenied())
	} else if title := r.FormValue("title"); title == "" {
		sendResponse(w, models.GetResponseErrNoValueInBody())
	} else {
		data := models.GetData()
		for i := 0; i < len(data.Items); i++ {
			if data.Items[i].Title == title {
				data.Items = append(data.Items[:i], data.Items[i+1:]...)
				i--
			}
		} // удаляем все (вдруг юзер идиот) объекты, с указанным ключом (именем).
		data.Write()
		sendResponse(w, models.GetResponseOK())
	}
}
