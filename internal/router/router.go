package router

import "github.com/labstack/echo/v4"

type Router struct {
	engine   *echo.Echo
	listener string
}

const defaultListener = ":8080"

func New(options ...Option) (*Router, error) {
	r := &Router{
		engine:   echo.New(),
		listener: defaultListener,
	}

	for _, o := range options {
		o(r)
	}

	return r, nil
}

func (r *Router) Run() error {
	return r.engine.Start(r.listener)
}
