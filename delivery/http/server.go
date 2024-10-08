package http

import (
	"challenge/delivery/http/manager"
	"challenge/docs"
	"challenge/service"
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
)

type Config struct {
	Port int
}
type Server struct {
	config         Config
	echo           *echo.Echo
	managerHandler manager.Handler
}

func New(
	config Config,
	svc *service.Service,
) *Server {
	server := &Server{
		echo:           echo.New(),
		config:         config,
		managerHandler: manager.New(svc.ManagerSvc),
	}
	server.registerRoutes()
	return server
}

func (s Server) Serve() {
	address := fmt.Sprintf(":%d", s.config.Port)
	if err := s.echo.Start(address); err != nil {
		log.Fatalln(err)
	}
}

func (s Server) registerRoutes() {

	docs.SwaggerInfo.Title = "Challenge API"
	docs.SwaggerInfo.Description = "This is the API documentation for the Challenge project"
	docs.SwaggerInfo.Version = "1.0.0"
	s.echo.GET("/swagger/*any", echoSwagger.WrapHandler)

	s.managerHandler.SetRoutes(s.echo)
}
