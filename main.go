package main

import (
	"flag"
	"tacit/config"
	"tacit/server"
)

const (
	DEFAULT_CONFIG_FILE_PATH = "./config.yml"
	DEFAULT_PORT             = "8080"
)

func main() {

	// parse arguments
	var configFilePath string
	var port string
	flag.StringVar(&configFilePath, "f", "", "Config file path")
	flag.Parse()
	if configFilePath == "" {
		configFilePath = DEFAULT_CONFIG_FILE_PATH
	}
	if port == "" {
		port = DEFAULT_PORT
	}

	// get tacit config
	config, err := config.Read(configFilePath)
	if err != nil {
		panic(err)
	}

	// start server
	s := server.New()
	s.RegisterEndpoints(config.Endpoints)
	err = s.Listen()
	if err != nil {
		panic(err)
	}
}
