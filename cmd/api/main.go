package main

import (
	"dev/projects/training_server_bd/internal/app/api"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	flag.Parse()
	log.Println("It works")
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println(err)
	}

	server := api.New(config)

	log.Fatal(server.Start())
}
