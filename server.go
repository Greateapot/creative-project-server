package main

/*
TODO:
1. on path /get request return list of files from path as items (title = filename)
2. add logs to file or/and stream (stdout/stdin/stderr)
3. Вернуть /shutdown. Я вспомнил зачем он нужен был ;)
*/

import (
	"fmt"
	"greateapot/creative_project_server/models"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
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
	default:
		http.NotFound(w, r)
	}
}

func main() {
	if err := http.ListenAndServe(
		fmt.Sprintf("%s:%d", models.LocalIp, models.Port),
		http.HandlerFunc(handler),
	); err != http.ErrServerClosed {
		panic(err)
	}
}

// go build -o bin/
// bin/creative_project_server.exe --local-ip=192.168.10.104
