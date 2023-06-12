package models

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Online struct {
	Online []string `json:"online"`
}

func getOnline(a int, b int, pattern string, timeout time.Duration, buf chan int, ex int) {
	for i := a; i < b; i++ {
		if i == ex {
			continue // пропускаем себя
		}
		address := fmt.Sprintf("%s.%d:%d", pattern, i, Port)
		if _, err := net.DialTimeout("tcp", address, timeout); err == nil {
			buf <- i
		}
	}
	buf <- -1 // сигнал о том, что перебор окончен
}

func GetOnline() *Online {
	lips := strings.Split(LocalIp, ".")
	pattern := strings.Join(lips[:3], ".")
	ex, _ := strconv.Atoi(lips[3])
	buf := make(chan int)
	count := 0
	line := ""

	timeout := time.Duration(time.Duration(scanTimeout) * time.Millisecond)

	for a := 0; a < 256; a += 256 / scanThreads {
		go getOnline(a, a+(256/scanThreads), pattern, timeout, buf, ex)
	}

	for {
		v, ok := <-buf
		if !ok {
			break
		} else if v < 0 {
			count++
			if count == scanThreads {
				close(buf)
			}
		} else {
			line += "," + strconv.Itoa(v)
		}
	}
	online := strings.Split(strings.Trim(line, ","), ",")
	filtered := []string{}

	for i := range online {
		if len(online[i]) > 0 {
			filtered = append(filtered, pattern+"."+online[i])
		}
	}
	return &Online{Online: filtered}
}
