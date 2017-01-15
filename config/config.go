package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

var Config = struct {
	DB struct {
		Conn string `json:"Conn"`
	} `json:"DB"`
}{}

func InitConfig(cfg string) {
	// configor.Load(&Config, "config/config.json")
	 configor.Load(&Config, cfg)
	fmt.Println(Config.DB.Conn)
}
