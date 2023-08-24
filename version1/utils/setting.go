package utils

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	AppMode    string
	HttpPort   string
	JwtKey     string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误：", err)
	}
	LoadServer(file)
	LoadDataBase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":6000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("JsonWebTokens")
}

func LoadDataBase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbName = file.Section("database").Key("DbName").MustString("3306")
	DbPort = file.Section("database").Key("DbPort").MustString("root")
	DbUser = file.Section("database").Key("DbUser").MustString("wyc123..")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("rs")
}
