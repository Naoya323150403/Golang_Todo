package config

import (
	"Golang6/src/Todo_App/utils"
	"gopkg.in/ini.v1"
	"log"
)


type ConfigList struct{
	Port      string
	SQLDriver string
	DBName    string
	Logfile   string
	Static    string
}

var Config ConfigList

func init(){
	LoadConfig()
	utils.LoggingSetting(Config.Logfile)
}

func LoadConfig(){
	cfg, err := ini.Load("config.ini")
	if err != nil{
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DBName:    cfg.Section("db").Key("name").String(),
		Logfile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),

	}




}