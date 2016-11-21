package Foundation

import (
	"github.com/dapsframework/daps/Component"
	"github.com/dapsframework/daps/Foundation/Loader"
	_ "path"
	"path"
)

type Application struct {
	Component.Component

	Version 	string
	BasePath 	string
	EnvironmentPath string
	EnvironmentFile string
	Loaded		bool
	Components	[]Component.ComponentInterface
}

func (application *Application) InitComponent() {
	application.Version = "0.0.1"
}

func (application *Application) AddComponent(component Component.ComponentInterface) {
	application.Components = append(application.Components, component)

	component.InitComponent()
}

func (application *Application) GetComponents() []Component.ComponentInterface {
	return application.Components
}

func (application *Application) LoadConfig() {
	configLoader := new(Loader.AppConfigLoader)

	Component.BootComponent(configLoader)
	application.AddComponent(configLoader)

	configLoader.Load(path.Join(application.BasePath, "Config", "app.yaml"))
}

func NewApplication(basePath string) *Application {
	application := new(Application)
	application.BasePath = basePath

	application.LoadConfig()

	return application
}
