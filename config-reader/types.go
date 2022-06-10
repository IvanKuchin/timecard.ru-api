package config_reader

type Config struct {
	Listenport int
	Serverhost string
	Serverport int
}

type ConfigReader interface {
	GetConfig() (*Config, error)
}
