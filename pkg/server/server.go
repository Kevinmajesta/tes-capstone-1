package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Kevinmajesta/tes-capstone-1/pkg/response"
	"github.com/Kevinmajesta/tes-capstone-1/pkg/route"
	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
}

func NewServer(serverName string, publicRoutes, privateRoutes []*route.Route) *Server {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Hello, World!", nil))
	})

	v1 := e.Group(fmt.Sprintf("/%s/api/v1", serverName))

	if len(publicRoutes) > 0 {
		for _, v := range publicRoutes {
			v1.Add(v.Method, v.Path, v.Handler)
		}
	}

	if len(privateRoutes) > 0 {
		for _, v := range privateRoutes {
			v1.Add(v.Method, v.Path, v.Handler)
		}
	}

	return &Server{e}
}

func (s *Server) Run() {
	runServer(s)
	gracefulShutdown(s)
}

func runServer(srv *Server) {
	go func() {
		err := srv.Start(":8080")
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func gracefulShutdown(srv *Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal("Server Shutdown:", err)
		}
	}()
}
