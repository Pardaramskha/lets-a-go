package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var serverPort = ":5000"

func main() {
	fmt.Println("Starting API...")
	launchServer()
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
	router.Mount("/todos", TodoRoutes())

	fmt.Println("API listening to", serverPort)
	// same reason for not handling error
	http.ListenAndServe(serverPort, router)
	// no logging past this point
}

// TodoRoutes Function that returns a whole router for todos
func TodoRoutes() chi.Router {
	router := chi.NewRouter()
	todoHandler := TodoHandler{}

	router.Get("/", makeHTTPHandleFunc(todoHandler.GetTodos))
	router.Post("/", makeHTTPHandleFunc(todoHandler.CreateTodo))
	router.Get("/{id}", makeHTTPHandleFunc(todoHandler.GetTodo))
	router.Put("/{id}", makeHTTPHandleFunc(todoHandler.EditTodo))
	router.Delete("/{id}", makeHTTPHandleFunc(todoHandler.DeleteTodo))

	return router
}
