package main

import (
	"encoding/base64"
	"encoding/json"
	"greateapot/creative_project_server/models"
	// "net"
	"net/http"
	"strings"
)

func sendResponse(w http.ResponseWriter, r *models.Response) {
	if data, err := json.MarshalIndent(r, "", "  "); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(data))
	} else {
		// никогда не вернется...
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// stackoverflow
// func GetLocalIP() string {
// 	conn, _ := net.Dial("udp", "8.8.8.8:80") // как оказалось, самый быстрый способ узнать свой локальный айпи
// 	defer conn.Close()

// 	localAddr := conn.LocalAddr().(*net.UDPAddr)

// 	return strings.Split(localAddr.String(), ":")[0] // порт нам не интересен, udp же
// }

// stackoverflow
func DecodeB64(data string) string {
	//.Replace('+', '.').Replace('/', '_').Replace('=', '-');
	data = strings.Replace(data, ".", "+", -1)
	data = strings.Replace(data, "_", "/", -1)
	data = strings.Replace(data, "-", "=", -1)
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	base64.StdEncoding.Decode(base64Text, []byte(data))
	return strings.Trim(string(base64Text), string(byte(0)))
	// А всё, А EncodeB64 не будет
}
