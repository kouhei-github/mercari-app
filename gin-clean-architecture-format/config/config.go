package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

type ConfigList struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	SpreadID   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini") // configファイル読み込み
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1) // プログラム終了
	}

	Config = ConfigList{
		DbHost:     cfg.Section("db").Key("host").String(),
		DbPort:     cfg.Section("db").Key("port").String(),
		DbUser:     cfg.Section("db").Key("user").String(),
		DbPassword: cfg.Section("db").Key("password").String(),
		DbName:     cfg.Section("db").Key("name").String(),
		SpreadID:   cfg.Section("spreadsheet").Key("spreadsheetID").String(),
	}
}
