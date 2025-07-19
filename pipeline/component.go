package pipeline

import "net/http"

type ComponentContext struct {
	*http.Request
	http.ResponseWriter
	error
}

func (mvc *ComponentContext) Error(err error) {
	mvc.error = err
}

func (mvc *ComponentContext) GetError() error {
	return mvc.error
}

type MiddlewareComponent interface {
	Init()
	ProcessRequest(context *ComponentContext, next func(*ComponentContext))
}

type ServicesMiddlewareComponent interface {
	Init()
	ImplementsProcessRequestWithServices()
}
