package Component

type ComponentInterface interface {
	InitComponent()
}

type Component struct {
	IsBooted	bool
}

func BootComponent(component ComponentInterface) {
	component.InitComponent()
}
