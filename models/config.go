package models

/*
На самом деле это не конфиг, просто 4 из 6 полей нужны здесь, так что почему бы не парсить аргументы здесь?
*/

import "flag"

/*
 ScanDelay: (4 threads)
 64 * 1000ms = 64 sec (slow LAN)

 64 * 500ms = 32 sec (medium LAN)

 64 * 100ms = 6.4 sec (fast LAN)
*/

var (
	LocalIp       string
	corrupted     string
	itemsFileName string
	Port          int
	scanTimeout   int // ms
	scanThreads   int // 256 / scanThreads, must be 2^N
)

func init() {
	flag.StringVar(&LocalIp, "local-ip", "", "local ip")
	flag.StringVar(&corrupted, "corr-file-ext", ".crp", "corrupted filename extension")
	flag.StringVar(&itemsFileName, "data-filename", "data.json", "data filename")

	flag.IntVar(&Port, "port", 8097, "port")
	flag.IntVar(&scanTimeout, "scan-timeout", 500, "scan timeout")
	flag.IntVar(&scanThreads, "scan-threads", 4, "scan threads count")

	flag.Parse()

	if LocalIp == "" {
		panic("No local ip provided!")
	}
}
