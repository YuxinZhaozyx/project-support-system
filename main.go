package main

import (
	"flag"
	"fmt"
	"project-support-system/config"
	"project-support-system/server"
)

var (
	address    string
	configFile string
)

func parseFlag() {
	flag.StringVar(&address, "address", "127.0.0.1:8080", "the address like 127.0.0.1:8080, defaults to 127.0.0.1:8080")
	flag.StringVar(&configFile, "config_file", "config.yaml", "the config file of yaml format")
	flag.Parse()
}

func main() {
	parseFlag()

	err := config.Init(configFile)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	server.Run(address)
}
