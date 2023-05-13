package main

import (
	"fmt"
	"greateapot/creative_project_server/models"
	"net/http"
)

// creative_project_server.exe --lip=192.168.10.104 --port=8097 --sd=100 --st=8
func main() {
	server := &http.Server{}
	server.Addr = fmt.Sprintf("%s:%d", models.LocalIp, models.Port)
	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch path := r.URL.Path; path {
		case "/get": // http://192.168.XXX.XXX:8097/get?title={B64ENC}
			HandleGet(w, r)
		case "/list": // http://192.168.XXX.XXX:8097/list
			HandleList(w, r)
		case "/add": // http://192.168.XXX.XXX:8097/add?title={B64ENC}&path={B64ENC}&type=1 // lip-lock
			HandleAdd(w, r)
		case "/del": // http://192.168.XXX.XXX:8097/del?title={B64ENC} // lip-lock
			HandleDel(w, r)
		case "/shutdown": // http://192.168.XXX.XXX:8097/shutdown // lip-lock
			HandleShutdown(server, w, r)
		case "/online": // http://192.168.XXX.XXX:8097/online // lip-lock
			HandleOnline(w, r)
		default: // http://192.168.XXX.XXX:8097/*
			http.NotFound(w, r)
		}
	})
	// lip-lock - блокировка доступа с других устройств. Делать такие запросы может только хост.
	// порт 8097 - порт по умолчанию. Его можно изменить, если какой-то другой софт уже занял его.

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
