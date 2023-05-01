package main

import (
	"fmt"
	"greateapot/creative-project-server/models"
	"io"
	"net/http"
	"os"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("title")
	if title == "" {
		responseString(w, http.StatusBadRequest, "no title")
		return
	}

	data := models.GetData()

	var item models.Item
	for _, it := range data.Items {
		if it.Title == title {
			item = it
			break
		}
	}

	if item.Path == "" {
		responseString(w, http.StatusNotFound, "not found")
		return
	}

	path := DecodeB64(item.Path)
	file, err := os.Open(path)

	defer func() {
		file.Close()
	}()

	if err != nil {
		responseString(w, http.StatusNotFound, "not found")
		return
	}

	stat, _ := file.Stat()

	w.Header().Set("Content-Disposition", "attachment; filename="+stat.Name())
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))

	// stream the file to the client without fully loading it into memory
	io.Copy(w, file)
}
