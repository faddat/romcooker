package main

import (
	"flag"
	"fmt"
	"github.com/olebedev/config"
)

type Configuration interface {
	ParseYamlFile(filename string) (*config.Config, error)
}

/*
Initialize Configuration
Defines some configuration options in a global variable
*/

var CFG *config.Config

func loadConfig() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "The path to configuration")
	flag.Parse()
	if configPath == "" {
		panic("No config specified")
	}
	config, err := config.ParseYamlFile(configPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("config loaded from " + configPath)
	CFG = config.Flag()
}
