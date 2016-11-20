package Loader

import (
	"github.com/dapsframework/daps/Component"
	"path"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

type AppConfig struct {
	Version int8 `yaml:"version"`
	Name string `yaml:"name"`
}

type ConfigLoader struct {
	Component.Component

	Version 	string
	ConfigPath 	string
}

func (configLoader *ConfigLoader) InitComponent() {
	configLoader.Version = "0.0.1"

	configLoader.loadAppConfig()
}

func (configLoader *ConfigLoader) Load(basePath string) {
	configLoader.ConfigPath = path.Join(basePath, "Config")
}

func (configLoader *ConfigLoader) loadAppConfig() {
	configFile := path.Join(configLoader.ConfigPath, "app.yaml")

	configLoader.loadConfigFile(configFile)
}

func (configLoader *ConfigLoader) loadConfigFile(configFile string) {
	yamlFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	var config AppConfig

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Value: %#v\n", config.Name)
}


