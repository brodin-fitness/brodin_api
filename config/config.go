package config

import "os"

type appConfig struct {
	Port string
}

type config struct {
	App appConfig
}

var Values *config

// Assign environmental variables to the config struct
func Configure() *config {
	return &config{
		App: appConfig{
			Port: getEnv("PORT", "8080"),
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
