package main

import (
	"clean_architecture/internal/handler"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	authService "clean_architecture/internal/service/login/deivery/http"
)

func main() {
	startHttpServer()
}

func startHttpServer() {
	initGinHttp()
}

func initGinHttp() {
	r := gin.Default()
	rootHandler := handler.NewRootHandler()
	authService.InitRoutes(r, rootHandler)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
