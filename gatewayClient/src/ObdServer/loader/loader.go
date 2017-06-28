package loader

import (
	"config"
)

func GetConfig() *ObdServerConfig {
	var v = newObdServerConfig()
	//config.GetConfig("src/simpleServer/cnf/cnf.json", v)
	config.GetConfig("../simpleServer/cnf/cnf.json", v)
	return v
}
