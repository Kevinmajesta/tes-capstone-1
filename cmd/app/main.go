package main

import (
	"github.com/Kevinmajesta/tes-capstone-1/configs"
	"github.com/Kevinmajesta/tes-capstone-1/internal/http/router"
	"github.com/Kevinmajesta/tes-capstone-1/pkg/server"
)

func main() {
	// Memuat konfigurasi dari file .env
	_, err := configs.NewConfig(".env")
	checkError(err)

	// Mengambil rute publik dan privat
	publicRoutes := router.PublicRoutes()
	privateRoutes := router.PrivateRoutes()

	// Membuat server baru dengan rute-rute yang telah diambil
	srv := server.NewServer("app", publicRoutes, privateRoutes)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
