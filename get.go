package main

import (
	"fmt"
	"greateapot/creative-project-server/models"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
TODO: в зависимости от типа:
1: уже готово (мб стоит изменить возврат пути хосту на формат file:///{...}) (ключ: title)
2: прверять, предоставляет ли хост доступ к этой папке, если да, дать запрошенный файл, нет - развернуть на 180 и
отправить обратно. (ключ: title+path)
3: вернуть, как если бы хост хотел получить свой файл (ключ: title)
*/
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

	if strings.Split(r.RemoteAddr, ":")[0] == local_ip {
		responseString(w, http.StatusOK, path) // Ну а зачем мне скачивать файл, который находится на моем устройстве?
		return                                 // Логичнее ж получить его путь и открыть в проводнике, чем заново скачивать.
	}

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
