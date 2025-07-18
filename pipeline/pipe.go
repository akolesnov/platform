package pipeline

import "net/http"

type RequestPipeline func(*ComponentContext)

var emptyPipeline RequestPipeline = func(*ComponentContext) {
	/* do nothing */
}

func CreatePipeline(components ...MiddlewareComponent) RequestPipeline {
	f := emptyPipeline
	for i := 0; i < len(components); i++ {
		currentComponent := components[i]
		nextFunc := f
		f = func(context *ComponentContext) {
			if context.error == nil {
				currentComponent.ProcessRequest(context, nextFunc)
			}
		}
		currentComponent.Init()
	}
	return f
}

func (pl RequestPipeline) ProcessRequest(req *http.Request, resp http.ResponseWriter) error {
	ctx := ComponentContext{
		Request:        req,
		ResponseWriter: resp,
	}
	pl(&ctx)
	return ctx.error
}
