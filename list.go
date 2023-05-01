package main

import (
	"fmt"
	"greateapot/creative-project-server/models"
	"net/http"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	data := models.GetData()
	result := ""
	for _, item := range data.Items {
		result += fmt.Sprintf("%d", item.Type) + ":" + item.Title + "\n"
	}
	responseString(w, http.StatusOK, result)
}
