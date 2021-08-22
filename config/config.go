package config

import "os"

type AppConfig struct {
	Port string
}

type Config struct {
	App AppConfig
}

func New() *Config {
	return &Config{
		App: AppConfig{
			Port: getEnv("PORT", ":8080"),
		},
	}
}

// Searches the environment for a variable, returns value if found or default if not found.
func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}