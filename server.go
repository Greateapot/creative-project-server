package main

/*
TODO:
1. add logs to file or/and stream (stdout/stdin/stderr)
*/

import (
	"fmt"
	"greateapot/creative_project_server/models"
	"net/http"
)

// go build -o bin/
// bin/creative_project_server.exe --local-ip=192.168.10.104
func main() {
	server := &http.Server{}
	server.Addr = fmt.Sprintf("%s:%d", models.LocalIp, models.Port)
	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// lip-lock - блокировка доступа с других устройств. Делать такие запросы может только хост.
		switch path := r.URL.Path; path {
		case "/get":
			HandleGet(w, r)
		case "/list":
			HandleList(w, r)
		case "/add":
			HandleAdd(w, r) // lip-lock
		case "/del":
			HandleDel(w, r) // lip-lock
		case "/online":
			HandleOnline(w, r) // lip-lock
		case "/shutdown":
			HandleShutdown(server, w, r) // lip-lock
		default:
			http.NotFound(w, r)
		}
	})

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
}
