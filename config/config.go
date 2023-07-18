package config

// Config is the global config for the application
type Config struct {
	Host        string
	Port        string
	DatabaseURL string
}

var config *Config

// Init initializes the config package
func Init() {
	config = &Config{
		Host:        LookupEnv("HOST", "0.0.0.0"),
		Port:        LookupEnv("PORT", "9000"),
		DatabaseURL: LookupEnv("DATABASE_URL", ""),
	}
}

// GetConfig returns the global config
func GetConfig() *Config { return config }

// Host returns the host
func Host() string { return config.Host }

// Port returns the port
func Port() string { return config.Port }

// DatabaseURL returns the database url
func DatabaseURL() string { return config.DatabaseURL }
