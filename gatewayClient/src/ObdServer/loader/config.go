package loader

type ObdServerConfig struct {
	Server_ipandport string `json:"server_ipandport"`
	Client_num       int    `json:"client_num"`
}

func newObdServerConfig() *ObdServerConfig {
	return &ObdServerConfig{}
}
