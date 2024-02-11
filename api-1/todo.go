package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type TodoHandler struct {
}

// Handlers
// ---------------------------------------------------------------------------------------------------------------------
func (handler TodoHandler) GetTodos(responseWriter http.ResponseWriter, request *http.Request) error {
	err := WriteJSON(responseWriter, http.StatusOK, listTodos())
	if err != nil {
		// Returns a 500
		http.Error(responseWriter, "Internal server error", http.StatusInternalServerError)
		return err
	}

	return nil
}

func (handler TodoHandler) GetTodo(responseWriter http.ResponseWriter, request *http.Request) error {
	id := chi.URLParam(request, "id")
	todo := getTodo(id)

	if todo == nil {
		http.Error(responseWriter, "Todo not found", http.StatusNotFound)
	}

	err := WriteJSON(responseWriter, http.StatusOK, todo)
	if err != nil {
		http.Error(responseWriter, "Server error", http.StatusInternalServerError)
		return err
	}

	return nil

}
func (handler TodoHandler) CreateTodo(responseWriter http.ResponseWriter, request *http.Request) error {
	var todo Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		http.Error(responseWriter, "Decoding error", http.StatusInternalServerError)
		return err
	}

	createTodo(todo)
	err = WriteJSON(responseWriter, http.StatusCreated, todo)
	if err != nil {
		http.Error(responseWriter, "Encoding error", http.StatusInternalServerError)
		return err
	}

	return nil
}
func (handler TodoHandler) EditTodo(responseWriter http.ResponseWriter, request *http.Request) error {
	id := chi.URLParam(request, "id")
	var todo Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		http.Error(responseWriter, "Decoding error", http.StatusInternalServerError)
		return err
	}

	updatedTodo := updateTodo(id, todo)
	if updatedTodo == nil {
		http.Error(responseWriter, "Todo not found", http.StatusNotFound)
	}

	err = WriteJSON(responseWriter, http.StatusOK, todo)
	if err != nil {
		http.Error(responseWriter, "Internal server error", http.StatusBadRequest)
		return err
	}

	return nil
}
func (handler TodoHandler) DeleteTodo(responseWriter http.ResponseWriter, request *http.Request) error {
	id := chi.URLParam(request, "id")
	todo := deleteTodo(id)
	if todo == nil {
		http.Error(responseWriter, "Todo not found", http.StatusNotFound)
	}
	responseWriter.WriteHeader(http.StatusNoContent)
	return nil
}

// Functions - part to modify following any DB system change
// ---------------------------------------------------------------------------------------------------------------------
func getTodo(id string) *Todo {
	// rolling in
	for _, todo := range todos {
		if todo.ID == id {
			return todo
		}
	}
	return nil
}

func createTodo(todo Todo) {
	todos = append(todos, &todo)
}

func deleteTodo(id string) *Todo {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], (todos)[i+1:]...)
			return &Todo{}
		}
	}
	return nil
}

func updateTodo(id string, update Todo) *Todo {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = &update
			return todo
		}
	}
	return nil
}
