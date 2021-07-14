package main

import (
	"encoding/json"
	"fmt"
	"os"
)


type setting struct {
	ServerHost string
	ServerPort string
	MySQLHost  string
	MySQLPort  string
	MySQLUser  string
	MySQLPass  string
	MySQLBase  string
	Data       string
	Assets     string
	HTML       string
}

var cfg setting

func init()  {
	//открыть файл конфигурации
	file, e := os.Open("setting.cfg")
	if e != nil{
		fmt.Println(e.Error())
		panic("Не удалось открыть файл конфишурации")
	}
	//отоженый вызов
	defer file.Close()

	//прочесть статистику
	stat, e := file.Stat()
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось прочесть информацию о файле конфигурации")
	}

	readByte := make([]byte, stat.Size())

	_, e = file.Read(readByte)
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось прочесть файл конфигурации")
	}

	e = json.Unmarshal(readByte, &cfg)
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось считать данных файла конфигурации")
	}

}
