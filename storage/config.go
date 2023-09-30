package storage

type Config struct {
	DatabaseURI string `toml:"databaseuri"`
}

func NewConfig() *Config {
	return &Config{
		
	}
}