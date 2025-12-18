package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiHandlers interface {
	RegisterRoutes(app *gin.Engine)
}

type Api struct {
	Gin      *gin.Engine
	Handlers []ApiHandlers
	Port     int
	Listen   string
}

func NewApi(listen string, port int) *Api {
	return &Api{
		Gin:      gin.Default(),
		Handlers: []ApiHandlers{},
		Port:     port,
		Listen:   listen,
	}
}

func (a *Api) AddHandler(handler ApiHandlers) {
	a.Handlers = append(a.Handlers, handler)
}

func (a *Api) Start() error {
	for _, handler := range a.Handlers {
		handler.RegisterRoutes(a.Gin)
	}
	return a.Gin.Run(fmt.Sprintf("%s:%d", a.Listen, a.Port))
}
