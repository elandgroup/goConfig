package init

import "goConfig/config"

func init() {
	config.InitConfig("./config/config.json")
}
