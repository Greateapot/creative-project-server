package main

import (
	"encoding/base64"
	"fmt"
	"greateapot/creative-project-server/models"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func responseString(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func GetLocalIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80") // как оказалось, самый быстрый способ узнать свой локальный айпи
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return strings.Split(localAddr.String(), ":")[0] // порт нам не интересен, udp же
}

func getOnline(a int, b int, pattern string, port string, timeout time.Duration, buf chan int, ex int) {
	for i := a; i < b; i++ {
		if i == ex {
			continue // пропускаем себя
		}
		if _, err := net.DialTimeout("tcp", pattern+"."+fmt.Sprintf("%d", i)+":"+port, timeout); err == nil {
			buf <- i
		}
	}
	buf <- -1 // сигнал о том, что перебор окончен
}

func GetOnline() string {
	lips := strings.Split(local_ip, ".") // Да, губы
	pattern := strings.Join(lips[:3], ".")
	ex, _ := strconv.Atoi(lips[3])
	buf := make(chan int)
	count := 0
	line := ""

	// cfg
	port := models.GetConfig().Port
	timeout := time.Duration(time.Millisecond * 500) // TODO: MV to cfg
	// 64 * 1 = 64 (slow LAN); 64 * .5 = 32 sec (medium LAN); 64 * .1 = 6.4 sec (fast LAN).

	for a := 0; a < 256; a += 64 {
		go getOnline(a, a+64, pattern, port, timeout, buf, ex)
	}

	for {
		v, ok := <-buf
		if !ok {
			break
		} else if v < 0 {
			count++
			if count == 4 {
				close(buf)
			}
		} else {
			line += "," + fmt.Sprintf("%d", v)
		}
	}

	return strings.Trim(line, ",") // лишняя ','
}

func DecodeB64(data string) string {
	//.Replace('+', '.').Replace('/', '_').Replace('=', '-');
	data = strings.Replace(data, ".", "+", -1)
	data = strings.Replace(data, "_", "/", -1)
	data = strings.Replace(data, "-", "=", -1)
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	base64.StdEncoding.Decode(base64Text, []byte(data))
	return strings.Trim(string(base64Text), string(byte(0)))
}
