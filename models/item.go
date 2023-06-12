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

type Items struct {
	Items []*Item `json:"items"`
}

func GetItems() *Items {
	Items := &Items{}
	Items.Read()
	return Items
}

func (d *Items) HidePath() (hiddenItems *Items) {
	hiddenItems = &Items{}
	for _, item := range d.Items {
		hiddenItem := &Item{item.Title, "", item.Type} // create copy with no path
		hiddenItems.Items = append(hiddenItems.Items, hiddenItem)
	}
	return
}

func (d *Items) open(flag int) (*os.File, error) {
	if homeDir, err := os.UserHomeDir(); err != nil {
		return nil, err
	} else {
		return os.OpenFile(filepath.Join(homeDir, "Documents", "Creative Project", itemsFileName), flag, 0666)
	}
}

func (d *Items) Read() {
	file, err := d.open(os.O_RDONLY)

	if err != nil {
		fmt.Println("Read(Open):", err, ";Items:", d)
		if !os.IsNotExist(err) {
			// Что-то не так с файлом
			file.Close()
			os.Rename(itemsFileName, itemsFileName+corrupted)
		}
		d.Write() // создаем новый
	} else {
		if buf, err := io.ReadAll(file); err != nil {
			fmt.Println("Read(ReadAll):", err, ";Items:", d)
		} else if err := json.Unmarshal(buf, d); err != nil {
			fmt.Println("Read(Unmarshal):", err, ";Items:", d)
			file.Close()
			os.Rename(itemsFileName, itemsFileName+corrupted)
			d.Write()
		} else {
			// fmt.Println("Read(End): ok")
			file.Close()
		}
	}
}

func (d *Items) Write() {
	file, err := d.open(os.O_CREATE | os.O_WRONLY | os.O_TRUNC)

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Write(Close):", err, ";Items:", d)
		}
	}()

	if err != nil {
		fmt.Println("Write(Open):", err, ";Items:", d)
	} else if Items, err := json.Marshal(d); err != nil {
		fmt.Println("Write(Marshal):", err, ";Items:", d)
	} else if _, err := file.Write(Items); err != nil {
		fmt.Println("Write(Write):", err, ";Items:", d)
	}
	// else {
	// fmt.Println("Write(End): ok")
	// }
}
