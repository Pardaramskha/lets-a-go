package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/pardaramskha/lets-a-go/api-2/handlers"
	"github.com/pardaramskha/lets-a-go/api-2/utils"
)

// TodoRoutes Function that returns a whole router for todos
func TodoRoutes() chi.Router {
	router := chi.NewRouter()
	todoHandler := handlers.TodoHandler{}

	router.Get("/", utils.MakeHTTPHandleFunc(todoHandler.GetTodos))
	router.Get("/{id}", utils.MakeHTTPHandleFunc(todoHandler.GetTodo))
	//router.Post("/", makeHTTPHandleFunc(todoHandler.CreateTodo))
	//router.Get("/{id}", makeHTTPHandleFunc(todoHandler.GetTodo))
	//router.Put("/{id}", makeHTTPHandleFunc(todoHandler.EditTodo))
	//router.Delete("/{id}", makeHTTPHandleFunc(todoHandler.DeleteTodo))

	return router
}
