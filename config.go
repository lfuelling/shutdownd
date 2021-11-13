package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	AuthUsername       string
	AuthPassword       string
	ListenAddress      string
	OsType             string
	UseSudo            bool
	UseTls             bool
	TlsCertificateFile string
	TlsCertificateKey  string
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
