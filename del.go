package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
)

func HandleDel(w http.ResponseWriter, r *http.Request) {
	if !IsHostRequest(r) {
		sendResponse(w, http.StatusForbidden, models.ErrAccessDenied())
	} else if title := r.FormValue("title"); title == "" {
		sendResponse(w, http.StatusBadRequest, models.ErrNoRequiredFieldInRequestQuery())
	} else {
		items := models.GetItems()
		for i := 0; i < len(items.Items); i++ {
			if items.Items[i].Title == title {
				items.Items = append(items.Items[:i], items.Items[i+1:]...)
				i--
			}
		} // удаляем все (вдруг юзер идиот) объекты, с указанным ключом (именем).
		items.Write()
		sendResponse(w, http.StatusOK, nil)
	}
}
