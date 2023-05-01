package main

import (
	"greateapot/creative-project-server/models"
	"net/http"
	"strconv"
	"strings"
)

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] != local_ip {
		responseString(w, http.StatusForbidden, "access denied")
		return
	}

	item := &models.Item{}

	item.Path = r.FormValue("path")
	if item.Path == "" {
		responseString(w, http.StatusBadRequest, "no path")
		return
	}

	if r.FormValue("type") == "" {
		responseString(w, http.StatusBadRequest, "no type")
		return
	}

	if t, err := strconv.ParseUint(r.FormValue("type"), 10, 8); err == nil {
		item.Type = uint8(t)
	} else {
		responseString(w, http.StatusBadRequest, "err type")
		return
	}

	item.Title = r.FormValue("title")
	if item.Title == "" {
		responseString(w, http.StatusBadRequest, "no title")
		return
	}

	data := models.GetData()

	flag := false
	for i := 0; !flag && i < len(data.Items); i++ {
		flag = data.Items[i].Title == item.Title
	}

	if flag {
		responseString(w, http.StatusBadRequest, "title already exists")
	} else {
		data.Items = append(data.Items, *item)
		data.Write()
		responseString(w, http.StatusOK, "ok")
	}
}
