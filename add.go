package main

import (
	"greateapot/creative_project_server/models"
	"net/http"
	"strconv"
)

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	if !IsHostRequest(r) {
		sendResponse(w, http.StatusForbidden, models.ErrAccessDenied())
		return
	}

	if title := r.FormValue("title"); title == "" {
		sendResponse(w, http.StatusBadRequest, models.ErrNoRequiredFieldInRequestQuery()) // TODO: redo: no title provided
	} else if path := r.FormValue("path"); path == "" {
		sendResponse(w, http.StatusBadRequest, models.ErrNoRequiredFieldInRequestQuery()) // TODO: redo: no path provided
	} else if itemType, err := strconv.Atoi(r.FormValue("type")); err != nil {
		sendResponse(w, http.StatusBadRequest, models.ErrNoRequiredFieldInRequestQuery()) // TODO: redo: cant parse
	} else if itemType < 1 || itemType > 3 {
		sendResponse(w, http.StatusBadRequest, models.ErrNoRequiredFieldInRequestQuery()) // TODO: redo: no type $itemType
	} else {
		item := &models.Item{
			Title: title,
			Path:  path,
			Type:  itemType,
		}

		items := models.GetItems()

		flag := false
		for i := 0; !flag && i < len(items.Items); i++ {
			flag = items.Items[i].Title == item.Title
		} // Ищем уже существующий

		if flag {
			sendResponse(w, http.StatusBadRequest, models.ErrItemAlreadyExists())
		} else {
			items.Items = append(items.Items, item)
			items.Write()
			sendResponse(w, http.StatusOK, nil)
		}
	}
}
