package apiserver

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var apiServer *APIServer

type APIServer struct {
	ginInstance *gin.Engine
}

func NewApiServer() *APIServer {

	ginInstance := gin.New()

	ginInstance.Use(gin.Recovery())

	ginInstance.Use(cors.Default())

	registerRoutes(ginInstance)

	apiServer = &APIServer{
		ginInstance: ginInstance,
	}
	return apiServer
}

func (s *APIServer) Start() {
	url := fmt.Sprintf(":%d", 3000)
	if err := s.ginInstance.Run(url); err != nil {
		log.Fatal("can't start api server on", url)
	}
}

func registerRoutes(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.Data(200, "text/html; charset=utf-8", []byte("Welcome to ATcnea's API"))
	})

	// v1 := r.Group("/v1")

}
