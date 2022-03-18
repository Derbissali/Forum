package server

import (
	"encoding/json"
	"fmt"
	"os"
)

type setting struct {
	ServerHost string
	ServerPort string
	Assets     string
	HTML       string
}

var cfg setting

func init() {
	file, e := os.Open("setting.cfg")
	if e != nil {
		fmt.Println(e)
		return
	}
	defer file.Close()

	stat, e := file.Stat()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	readByte := make([]byte, stat.Size())

	_, e = file.Read(readByte)
	if e != nil {
		fmt.Println(e)
		return
	}
	json.Unmarshal(readByte, &cfg)
}
