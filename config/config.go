package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	DB          DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

var configrations *Config

func loadConfig() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, falling back to system environment variables")
	}

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		fmt.Println("Http port missing !")
		os.Exit(1)
	}

	version := os.Getenv("Version")
	if version == "" {
		fmt.Println("Version missing !")
		os.Exit(1)
	}

	serviceName := os.Getenv("ServiceName")
	if serviceName == "" {
		fmt.Println("Service name missing !")
		os.Exit(1)
	}

	port, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("Invalid http port:", err)
		os.Exit(1)
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println("Invalid db port:", err)
		os.Exit(1)
	}

	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     dbPort,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DATABASE"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	if dbConfig.SSLMode == "" {
		dbConfig.SSLMode = "disable"
	}
	if dbConfig.Host == "" || dbConfig.User == "" || dbConfig.Name == "" {
		fmt.Println("Database config missing !")
		os.Exit(1)
	}

	configrations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
		DB:          dbConfig,
	}

}

func GetConfig() *Config {
	if configrations == nil {
		loadConfig()
	}
	return configrations
}
