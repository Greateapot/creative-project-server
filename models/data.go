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

Folder: 0

File: 1

Link: 2
*/
type Item struct {
	Path string `json:"path"`

	Type int `json:"type"`

	Title string `json:"title"`
}

type Data struct {
	Items []Item `json:"items"`
}

type HiddenItem struct {
	Type  int    `json:"type"`
	Title string `json:"title"`
}

type HiddenData struct {
	Items []HiddenItem `json:"items"`
}

func GetData() *Data {
	data := &Data{}
	data.Read()
	return data
}

func GetHiddenData() *HiddenData {
	data := &HiddenData{}
	data.Read()
	return data
}

func (d *Data) open(flag int) (*os.File, error) {
	if homeDir, err := os.UserHomeDir(); err != nil {
		return nil, err
	} else {
		return os.OpenFile(filepath.Join(homeDir, "Documents", "Creative Project", dataFileName), flag, 0666)
	}
}

func (d *HiddenData) open(flag int) (*os.File, error) {
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

func (d *HiddenData) Read() {
	file, err := d.open(os.O_RDONLY)

	if err != nil {
		fmt.Println("Read(Open):", err, ";data:", d)
		if !os.IsNotExist(err) {
			// Что-то не так с файлом
			file.Close()
			os.Rename(dataFileName, dataFileName+corrupted)
		}
		GetData().Write() // создаем новый
	} else {
		if buf, err := io.ReadAll(file); err != nil {
			fmt.Println("Read(ReadAll):", err, ";data:", d)
		} else if err := json.Unmarshal(buf, d); err != nil {
			fmt.Println("Read(Unmarshal):", err, ";data:", d)
			file.Close()
			os.Rename(dataFileName, dataFileName+corrupted)
			GetData().Write()
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
