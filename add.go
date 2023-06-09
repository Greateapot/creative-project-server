package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"strconv"
)

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	if !IsHostRequest(r) {
		sendResponse(w, models.GetResponseErrAccessDenied())
		return
	}

	if title := r.FormValue("title"); title == "" {
		sendResponse(w, models.GetResponseErrNoValueInBody()) // TODO: redo: no title provided
	} else if path := r.FormValue("path"); path == "" {
		sendResponse(w, models.GetResponseErrNoValueInBody()) // TODO: redo: no path provided
	} else if itemType, err := strconv.Atoi(r.FormValue("type")); err != nil {
		sendResponse(w, models.GetResponseErrNoValueInBody()) // TODO: redo: cant parse
	} else if itemType < 1 || itemType > 3 {
		sendResponse(w, models.GetResponseErrNoValueInBody()) // TODO: redo: no type $itemType
	} else {
		item := &models.Item{
			Title: title,
			Path:  path,
			Type:  itemType,
		}

		data := models.GetData()

		flag := false
		for i := 0; !flag && i < len(data.Items); i++ {
			flag = data.Items[i].Title == item.Title
		} // Ищем уже существующий

		if flag {
			sendResponse(w, models.GetResponseErrItemAlreadyExists())
		} else {
			data.Items = append(data.Items, item)
			data.Write()
			sendResponse(w, models.GetResponseOK())
		}
	}
}
