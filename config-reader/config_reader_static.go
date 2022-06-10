package config_reader

type StaticConfigReader struct {
}

func (scr *StaticConfigReader) GetConfig() (*Config, error) {
	return &Config{
		Listenport: 81,
		Serverhost: "127.0.0.1",
		Serverport: 80,
	}, nil
}

func NewStaticConfigReader() (ConfigReader, error) {
	return &StaticConfigReader{}, nil
}
