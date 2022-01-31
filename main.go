package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RolloCasanova/dispatch-workshop-4/controller"
	"github.com/RolloCasanova/dispatch-workshop-4/router"
	"github.com/RolloCasanova/dispatch-workshop-4/service/db"
	rd "github.com/RolloCasanova/dispatch-workshop-4/service/redis"
	"github.com/RolloCasanova/dispatch-workshop-4/usecase"

	"github.com/gorilla/handlers"
)

func main() {
	// create instances for the service, usecase, controller and router
	// injecting the corresponding dependencies to each one of them
	employeeDBService := db.New(nil)
	employeeRedisService := rd.New("", "", 0, 0)
	employeeUsecase := usecase.New(employeeDBService, employeeRedisService)
	employeeController := controller.New(employeeUsecase)
	httpRouter := router.Setup(employeeController)

	// Info to set up the server
	// don't use magic naming and magic numbers, there are better ways to do so (viper - covered in another workshop)
	host := "localhost"
	port := 8080

	// create http.Server instance
	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", host, port),
		Handler:           handlers.LoggingHandler(os.Stdout, httpRouter),
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	log.Printf("starting server in address, %s\n", server.Addr)
	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("starting server: %v", err)
	}
}
