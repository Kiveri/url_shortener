package config

type Config struct {
	DatabaseType string
}

func NewConfig(databaseType string) *Config {
	return &Config{
		DatabaseType: databaseType,
	}

}
