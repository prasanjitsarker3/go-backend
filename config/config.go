package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

  

type Config struct  {
	Version string
	ServiceName string
	HttpPort  int

}


var configrations * Config;


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

	configrations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
	}	

}


func GetConfig() * Config {
	if configrations ==nil {
		loadConfig()
	}
	return configrations
}	