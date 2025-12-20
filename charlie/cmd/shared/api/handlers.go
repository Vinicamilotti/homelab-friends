package api

import (
	"fmt"
	"log"
	"os"
	"strconv"

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

func NewApi() *Api {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil || port > 65535 {
		log.Printf("invalid PORT eviroment variable: '%s', using default 8082", os.Getenv("PORT"))
		port = 8082
	}

	return &Api{
		Gin:      gin.Default(),
		Handlers: []ApiHandlers{},
		Port:     port,
		Listen:   os.Getenv("HOST"),
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
