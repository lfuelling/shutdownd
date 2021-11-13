package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	authUsername  string
	authPassword  string
	listenAddress string
	osType        string
	useSudo       bool
}

func readConfig(configPath string) (Config, error) {
	log.Println("Loading config from '" + configPath + "'")
	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	config := Config{}
	decoder := json.NewDecoder(file)
	err1 := decoder.Decode(&config)
	return config, err1
}
