package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vincen320/auth-service-mongodb/controller"
	"github.com/vincen320/auth-service-mongodb/middleware"
	"github.com/vincen320/auth-service-mongodb/service"
)

func main() {
	validator := validator.New()
	AuthServices := service.NewAuthService(validator)
	AuthController := controller.NewAuthController(AuthServices)

	router := gin.New()
	router.Use(middleware.ErrorHandling())
	router.POST("/login", AuthController.Login)
	router.GET("/refresh", AuthController.Refresh)

	server := &http.Server{
		Addr:           ":8081",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Auth Service Start in 8081 port")
	err := server.ListenAndServe()
	if err != nil {
		panic("ERROR STARTING SERVER")
	}
	//BUAT PRODUCT-SERVICE
	//UNTUK BUAT PRODUCT HARUS DIAUTHORIZATION DULU DENGAN JWT
}
