package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Engine is server interface for Web server will hold all muxer , middlware and configuration
type Engine struct {
	Name   string
	Port   int
	Server *gin.Engine
}

func ApplicationServer(name string, port int) *Engine {
	fmt.Printf("Start Creating New Web Server for blocker")
	return &Engine{Name: name, Port: port}
}

func (engine Engine) RegisterRouter(method string, path string, handlers gin.HandlerFunc) {
	engine.Server.Handle(method, path, handlers)

}

func (engine Engine) Run() {
	engine.Server.Run(":" + strconv.Itoa(engine.Port))
}
