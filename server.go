package main

import (
	"greateapot/creative-project-server/models"
	"log"
	"net/http"
)

var (
	server   = &http.Server{} // чтоб хендлер handleShutdown мог находится в отдельном файле
	local_ip = ""             // если б локальный айпи не менялся каждую перезагрузку роутера...
)

func main() {
	local_ip = GetLocalIP()

	server.Addr = local_ip + ":" + models.GetConfig().Port
	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch path := r.URL.Path; path {
		case "/get": // http://192.168.10.103:8097/get?title={B64ENC}
			HandleGet(w, r)
		case "/list": // http://192.168.10.103:8097/list
			HandleList(w, r)
		case "/add": // http://192.168.10.103:8097/add?title={B64ENC}&path={B64ENC}&type=1 // lip-lock
			HandleAdd(w, r)
		case "/del": // http://192.168.10.103:8097/del?title={B64ENC} // lip-lock
			HandleDel(w, r)
		case "/shutdown": // http://192.168.10.103:8097/shutdown // lip-lock
			HandleShutdown(w, r)
		case "/online": // http://192.168.10.103:8097/online // lip-lock
			HandleOnline(w, r)
		case "/config": // http:192.168.10.103:8097/config?dataFileName={B64ENC}&port=8097 // lip-lock
			HandleConfig(w, r)
		default: // http://192.168.10.103:8097/*
			http.NotFound(w, r)
		}
	})

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Panicln(err)
	}
}
