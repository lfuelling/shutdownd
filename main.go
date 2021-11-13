package main

import (
	"flag"
	"log"
)

func main() {
	// initialize flags
	configPath := flag.String("config", "/etc/shutdownd/config.json", "path to the config file")
	flag.Parse()

	// load configuration
	config, err := readConfig(*configPath)
	if err != nil {
		log.Fatalln("Unable to read config!", err)
	}

	// start the server
	startServer(config)
}
