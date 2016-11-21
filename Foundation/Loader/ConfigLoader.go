package Loader

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/dapsframework/daps/Component"
)

type ConfigLoaderInterface interface {
	Load(configPath string)
}

type ConfigLoader struct {
	Component.Component

	Config map[interface{}]interface{}
}

func (configLoader *ConfigLoader) InitComponent() {
	configLoader.Config = make(map[interface{}]interface{})
}

func (configLoader *ConfigLoader) LoadConfigFile(configFile string) map[interface{}]interface{} {
	yamlFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	config := make(map[interface{}]interface{})

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		panic(err)
	}

	return config
}


