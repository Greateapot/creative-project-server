package models

import "flag"

/*
 ScanDelay: (4 threads)
 64 * 1000ms = 64 sec (slow LAN)

 64 * 500ms = 32 sec (medium LAN)

 64 * 100ms = 6.4 sec (fast LAN)
*/

var (
	LocalIp      string
	corrupted    string
	dataFileName string
	Port         int
	scanDelay    int // ms
	scanThreads  int // 256 / scanThreads, must be 2^N
)

func init() {
	flag.IntVar(&Port, "port", 8097, "port")
	flag.IntVar(&scanDelay, "sd", 500, "scan delay")
	flag.IntVar(&scanThreads, "st", 4, "scan threads count")

	flag.StringVar(&LocalIp, "lip", "", "local ip")
	flag.StringVar(&corrupted, "cfe", ".crp", "corrupted filename extension")
	flag.StringVar(&dataFileName, "df", "data.json", "data filename")

	flag.Parse()

	if LocalIp == "" {
		panic("No local ip provided!")
	}
}

// import (
// 	"encoding/json"
// 	"log"
// 	"os"
// )

// type Config struct {
// 	DataFileName string `json:"dataFileName"`
// 	Port         string `json:"port"`
// 	ScanDelay    int    `json:"scanDelay"`
// }

// const (
// 	ConfigFileName = "cfg.json"
// 	corrupted      = ".crp"
// )

// func GetConfig() *Config {
// 	config := &Config{"data.json", "8097", 500}
// 	config.Read()
// 	return config
// }

// func (c *Config) Read() {
// 	file, err := os.OpenFile(ConfigFileName, os.O_RDONLY, 0666)

// 	if err != nil {
// 		file.Close()
// 		c.Reset()
// 	} else {
// 		buf, _ := os.ReadFile(ConfigFileName)
// 		if err := json.Unmarshal(buf, c); err != nil {
// 			file.Close()
// 			c.Reset()
// 		} else {
// 			file.Close()
// 		}
// 	}
// }

// func (c *Config) Reset() {
// 	os.Remove(ConfigFileName)
// 	c.Write()
// }

// // TODO: RM logs
// func (c *Config) Write() {
// 	file, err := os.OpenFile(
// 		ConfigFileName,
// 		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
// 		0666,
// 	)

// 	defer func() {
// 		if err := file.Close(); err != nil {
// 			log.Panicln("Write(Close):", err, ";data:")
// 		}
// 	}()

// 	if err != nil {
// 		log.Panicln("Write(Open):", err, ";data:")
// 	}
// 	data, err := json.Marshal(c)
// 	if err != nil {
// 		log.Panicln("Write(Marshal):", err, ";data:")
// 	}
// 	if _, err := file.Write(data); err != nil {
// 		log.Panicln("Write(Write):", err, ";data:")
// 	}

// }
