package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

// App is a union of configs
type App struct {
	Logger
	Server
	MySQL
}

// Logger contains log config
type Logger struct {
	Level     string
	Output    string
	Formatter string
}

// Server contains server config
type Server struct {
	AppVersion string
	Port       string
	Mode       string
}

// MySQL contains MySQL config
type MySQL struct {
	Host         string
	Port         string
	User         string
	Password     string
	DBname       string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  time.Duration
}

// Get retrieves configs from config file and populates config structures above
func Get() (*App, error) {
	viper := viper.New()

	viper.SetConfigName(configPath())
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	var c App
	if err := viper.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &c, nil
}

func configPath() string {
	env := os.Getenv("config")
	switch env {
	case "docker":
		return "config/api_docker"
	}
	return "config/api_local"
}
