package Loader

import "path"

const APP_CONFIG_FILE = "app.yaml"

type AppConfig struct {
	Version int8 `yaml:"version"`
	Name string `yaml:"name"`
}

type AppConfigLoader struct {
	ConfigLoader
}

func (configLoader *AppConfigLoader) Load(configPath string) {
	config := configLoader.LoadConfigFile(path.Join(configPath, APP_CONFIG_FILE))

	configLoader.Config = config
}
