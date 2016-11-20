package Routing

import (
	"github.com/dpitkevics/daps/Middleware"
)

type Controller struct {
	middleware []Middleware.Middleware
}

func (controller *Controller) AddMiddleware(middleware Middleware.Middleware) {
	controller.middleware = append(controller.middleware, middleware)
}

func (controller *Controller) GetMiddleware() []Middleware.Middleware {
	return controller.middleware
}
