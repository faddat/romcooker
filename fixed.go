package main


import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"io/ioutil"
	"log"
	"github.com/olebedev/config"
)

type Configuration interface {
	ParseYamlFile(filename string) (*config.Config, error)
}
var CFG *config.Config

func loadConfig() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "The path to configuration")
	flag.Parse()
	if configPath == "" {
		Error.Panic("No config specified")
	}
	config, err := config.ParseYamlFile(configPath)
	if err != nil {
		Error.Println(err.Error())
		return
	}
	Info.Println("config loaded from " + configPath)
	CFG = config.Flag()
}



}

func ParseYamlFile(filename string) (*Config, error)