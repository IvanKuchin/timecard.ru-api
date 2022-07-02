package config_reader

func Read() (*Config, error) {
	config_reader, err := NewConfigReader("yaml")
	if err != nil {
		return nil, err
	}

	return config_reader.GetConfig()
}
