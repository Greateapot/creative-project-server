package main

import (
	"fmt"
	"greateapot/creative_project_server/models"
	"io"
	"net/http"
	"os"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	body, err := parseRequestBody(r)

	if err != nil {
		sendResponse(w, models.CreateErrResponse(2))
		return
	}

	if body.Title == "" {
		sendResponse(w, models.CreateErrResponse(4))
		return
	}

	data := models.GetData()

	var item models.Item
	for _, it := range data.Items {
		if it.Title == body.Title {
			item = it
			break
		}
	}

	if item.Type != 1 { // не файл
		sendResponse(w, models.CreateErrResponse(11))
		return
	}

	// TODO: RW
	// if strings.Split(r.RemoteAddr, ":")[0] == models.LocalIp {
	// 	/*
	// 		Ну а зачем мне скачивать файл, который находится на моем устройстве?
	// 		Логичнее ж получить его путь и открыть в проводнике, чем заново скачивать.
	// 	*/
	// 	sendResponse(w, models.CreateDataResponse(item.Path))
	// 	return
	// }

	file, err := os.Open(item.Path)

	defer func() {
		file.Close()
	}()

	if err != nil {
		sendResponse(w, models.CreateErrResponse(10))
		return
	}

	stat, _ := file.Stat()

	w.Header().Set("Content-Disposition", "attachment; filename="+stat.Name())
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))

	// stream the file to the client without fully loading it into memory
	io.Copy(w, file)
}
