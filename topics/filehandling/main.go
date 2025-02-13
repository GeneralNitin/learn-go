package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
)

//go:embed config.json
var configFile []byte

type Config struct {
	Props struct {
		LogLevel string `json:"log-level"`
		Port     string `json:"port"`
	} `json:"config"`
}

func main() {
	// Read the config file
	config, err := readConfig(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Print the config
	fmt.Printf("Log Level: %s\n", config.Props.LogLevel)
	fmt.Printf("Port: %s\n", config.Props.Port)
}

func readConfig(data []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
