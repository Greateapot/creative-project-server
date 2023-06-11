package main

import (
	"fmt"
	"greateapot/creative_project_server/models"
	"io"
	"net/http"
	"os"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {

	if title := r.FormValue("title"); title == "" {
		sendResponse(w, http.StatusBadRequest, models.GetResponseErrNoFieldInQuery())
	} else {
		var item *models.Item

		for _, i := range models.GetData().Items {
			if i.Title == title {
				item = i
				break
			}
		}
		if item.Title == "" {
			sendResponse(w, http.StatusBadRequest, models.GetResponseErrItemNotExists())
			return
		}

		switch item.Type {
		case 1: // folder
			handleFolder(w, item)
		case 2: // file
			if IsHostRequest(r) {
				// TODO: А зачем тебе твой же файл?
				sendResponse(w, http.StatusTeapot, models.GetResponseErr())
			} else {
				handleFile(w, item)
			}
		case 3: // link
			handleLink(w, item)
		default:
			// TODO: А как так вышло, что у нас в [data.json] объект с несуществующим типом?
			sendResponse(w, http.StatusTeapot, models.GetResponseErr())
		}
	}
}

func handleFolder(w http.ResponseWriter, item *models.Item) {
	sendResponse(w, http.StatusTeapot, models.GetResponseErrWIP())
}

func handleFile(w http.ResponseWriter, item *models.Item) {
	file, err := os.Open(item.Path)

	defer func() {
		file.Close()
	}()

	if err != nil {
		sendResponse(w, http.StatusInternalServerError, models.GetResponseErrSendFile())
	} else if stat, err := file.Stat(); err != nil {
		sendResponse(w, http.StatusInternalServerError, models.GetResponseErrSendFile())
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Disposition", "attachment; filename="+item.Title) // TODO: replace err symbols (\|'"...)
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
		io.Copy(w, file)
	}
}

func handleLink(w http.ResponseWriter, item *models.Item) {
	sendResponse(w, http.StatusOK, models.GetResponseData(item.Path))
}
