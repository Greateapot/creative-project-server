package models

import (
	"encoding/json"
	"log"
	"os"
)

type Item struct {
	Path  string `json:"path"`
	Type  byte   `json:"type"`
	Title string `json:"title"`
}

type Data struct {
	Items []Item `json:"items"`
}

func GetData() *Data {
	data := &Data{}
	data.Read()
	return data
}

func (d *Data) Read() {
	file, err := os.OpenFile(GetConfig().DataFileName, os.O_RDONLY, 0666)

	if err != nil {
		if !os.IsNotExist(err) {
			// Что-то не так с файлом
			file.Close()
			os.Rename(GetConfig().DataFileName, GetConfig().DataFileName+corrupted)
		}
		d.Write() // создаем новый
	} else {
		buf, _ := os.ReadFile(GetConfig().DataFileName)
		if err := json.Unmarshal(buf, d); err != nil {
			file.Close()
			os.Rename(GetConfig().DataFileName, GetConfig().DataFileName+corrupted)
			d.Write()
		} else {
			file.Close()
		}
	}
}

// TODO: RM logs
func (d *Data) Write() {
	file, err := os.OpenFile(
		GetConfig().DataFileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0666,
	)

	defer func() {
		if err := file.Close(); err != nil {
			log.Panicln("Write(Close):", err, ";data:", d)
		}
	}()

	if err != nil {
		log.Panicln("Write(Open):", err, ";data:", d)
	}
	data, err := json.Marshal(d)
	if err != nil {
		log.Panicln("Write(Marshal):", err, ";data:", d)
	}
	if _, err := file.Write(data); err != nil {
		log.Panicln("Write(Write):", err, ";data:", d)
	}

}
