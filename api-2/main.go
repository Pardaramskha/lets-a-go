package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/pardaramskha/lets-a-go/api-2/routes"
	"net/http"
	"os"
)

var serverPort string

func main() {
	fmt.Println("Starting API...")

	serverPort = setPort()
	launchServer()
}

func setPort() string {
	err := godotenv.Load()
	if err != nil {
		panic("No environment found!")
	}

	return ":" + os.Getenv("API_PORT")
}

func launchServer() {
	router := chi.NewRouter()
	// Logs information about incoming requests
	router.Use(middleware.Logger)

	// Router native list
	router.Get("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		// no need to handle error here - if it's fucked, the whole api is fucked, tl;dr
		responseWriter.Write([]byte("Healthcheck: OK"))
	})

	// Mount any set of routes on any given URL
	router.Mount("/todos", routes.TodoRoutes())

	fmt.Println("API listening to", serverPort)

	err := http.ListenAndServe(serverPort, router)
	if err != nil {
		// don't you forget to add the ":" before the port you bloody twat
		panic(err)
	}
	// no logging past this point
}
