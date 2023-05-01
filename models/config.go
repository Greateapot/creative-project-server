package models

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DataFileName string `json:"dataFileName"`
	Port         string `json:"port"`
}

const (
	ConfigFileName = "cfg.json"
	corrupted      = ".crp"
)

func GetConfig() *Config {
	config := &Config{"data.json", "8097"}
	config.Read()
	return config
}

func (c *Config) Read() {
	file, err := os.OpenFile(ConfigFileName, os.O_RDONLY, 0666)

	if err != nil {
		file.Close()
		c.Reset()
	} else {
		buf, _ := os.ReadFile(ConfigFileName)
		if err := json.Unmarshal(buf, c); err != nil {
			file.Close()
			c.Reset()
		} else {
			file.Close()
		}
	}
}

func (c *Config) Reset() {
	os.Remove(ConfigFileName)
	c.Write()
}

// TODO: RM logs
func (c *Config) Write() {
	file, err := os.OpenFile(
		ConfigFileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0666,
	)

	defer func() {
		if err := file.Close(); err != nil {
			log.Panicln("Write(Close):", err, ";data:")
		}
	}()

	if err != nil {
		log.Panicln("Write(Open):", err, ";data:")
	}
	data, err := json.Marshal(c)
	if err != nil {
		log.Panicln("Write(Marshal):", err, ";data:")
	}
	if _, err := file.Write(data); err != nil {
		log.Panicln("Write(Write):", err, ";data:")
	}

}
