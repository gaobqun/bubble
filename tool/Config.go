package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

// 	定义与配置一样的结构体
type Config struct {
	AppName  string   `json:"app_name"`
	AppMode  string   `json:"app_mode"`
	AppHost  string   `json:"app_host"`
	AppPort  string   `json:"app_port"`
	Database DataBase `json:"database"`
}
// 定义和数据库配置一样的结构体
type DataBase struct {
	Driver  string `json:"driver"`
	User    string `json:"user"`
	Pwd     string `json:"pwd"`
	Host    string `json:"host"`
	Port    string `json:"port"`
	DbName  string `json:"db_name"`
	Charset string `json:"charset"`
	Show    bool   `json:"show_sql"`
}

var _cfg *Config = nil

// 解析结构体
func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
