package config_reader

type Config struct {
	Listenport  int
	Serverproto string
	Serverhost  string
	Serverport  int
}

type ConfigReader interface {
	GetConfig() (*Config, error)
}
