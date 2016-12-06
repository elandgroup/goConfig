package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	DB struct {
		Conn string `json:"Conn"`
	} `json:"DB"`
}{}

func InitConfig() {
	configor.Load(&Config, "config/config.json")
}
