package main

import (
	"github.com/Kevinmajesta/tes-capstone-1/internal/builder"
	"github.com/Kevinmajesta/tes-capstone-1/pkg/server"
)

func main() {
	publicRoutes := builder.BuildAuthPublicRoutes()
	privateRoutes := builder.BuildAuthPrivateRoutes()

	srv := server.NewServer("auth", publicRoutes, privateRoutes)
	srv.Run()
}
