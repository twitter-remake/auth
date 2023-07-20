package config

// Config is the global config for the application
type Config struct {
	AppName     string
	Host        string
	Port        string
	DatabaseURL string
	Environment string
}

var config *Config

// Init initializes the config package
func Init() {
	config = &Config{
		AppName:     LookupEnv("APP_NAME", "auth-service"),
		Host:        LookupEnv("HOST", "0.0.0.0"),
		Port:        LookupEnv("PORT", "9000"),
		DatabaseURL: LookupEnv("DATABASE_URL", ""),
		Environment: LookupEnv("ENVIRONMENT", "development"),
	}
}

// GetConfig returns the global config
func GetConfig() *Config { return config }

// AppName returns the app name
func AppName() string { return config.AppName }

// Host returns the host
func Host() string { return config.Host }

// Port returns the port
func Port() string { return config.Port }

// DatabaseURL returns the database url
func DatabaseURL() string { return config.DatabaseURL }

// Environment returns the environment
func Environment() string { return config.Environment }
