package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/*
Item.Type:

Empty: 0 (default)

Folder: 1

File: 2

Link: 3
*/
type Item struct {
	Title string `json:"title"`

	Path string `json:"path,omitempty"`

	Type int `json:"type"`
}

type Data struct {
	Items []*Item `json:"items"`
}

func GetData() *Data {
	data := &Data{}
	data.Read()
	return data
}

func (d *Data) HidePath() (hiddenData *Data) {
	hiddenData = &Data{}
	for _, item := range d.Items {
		hiddenItem := &Item{item.Title, "", item.Type} // create copy with no path
		hiddenData.Items = append(hiddenData.Items, hiddenItem)
	}
	return
}

func (d *Data) open(flag int) (*os.File, error) {
	if homeDir, err := os.UserHomeDir(); err != nil {
		return nil, err
	} else {
		return os.OpenFile(filepath.Join(homeDir, "Documents", "Creative Project", dataFileName), flag, 0666)
	}
}

func (d *Data) Read() {
	file, err := d.open(os.O_RDONLY)

	if err != nil {
		fmt.Println("Read(Open):", err, ";data:", d)
		if !os.IsNotExist(err) {
			// Что-то не так с файлом
			file.Close()
			os.Rename(dataFileName, dataFileName+corrupted)
		}
		d.Write() // создаем новый
	} else {
		if buf, err := io.ReadAll(file); err != nil {
			fmt.Println("Read(ReadAll):", err, ";data:", d)
		} else if err := json.Unmarshal(buf, d); err != nil {
			fmt.Println("Read(Unmarshal):", err, ";data:", d)
			file.Close()
			os.Rename(dataFileName, dataFileName+corrupted)
			d.Write()
		} else {
			// fmt.Println("Read(End): ok")
			file.Close()
		}
	}
}

func (d *Data) Write() {
	file, err := d.open(os.O_CREATE | os.O_WRONLY | os.O_TRUNC)

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Write(Close):", err, ";data:", d)
		}
	}()

	if err != nil {
		fmt.Println("Write(Open):", err, ";data:", d)
	} else if data, err := json.Marshal(d); err != nil {
		fmt.Println("Write(Marshal):", err, ";data:", d)
	} else if _, err := file.Write(data); err != nil {
		fmt.Println("Write(Write):", err, ";data:", d)
	}
	// else {
	// fmt.Println("Write(End): ok")
	// }
}
