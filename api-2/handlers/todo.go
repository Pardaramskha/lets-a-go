package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/pardaramskha/lets-a-go/api-2/models"
	"github.com/pardaramskha/lets-a-go/api-2/storage"
	"github.com/pardaramskha/lets-a-go/api-2/utils"
)

type TodoHandler struct{}

// Handlers
// --------------------------------------------------------------------------------------------------------------------
func (handler TodoHandler) GetTodos(responseWriter http.ResponseWriter, request *http.Request) error {
	todos := getAllTodos()

	err := utils.WriteJSON(responseWriter, http.StatusOK, todos)
	if err != nil {
		log.Fatalf("Failed to write response: %v", err)
	} else {
		fmt.Println("Request completed: GET /todos")
	}

	return nil
}

func (handler TodoHandler) GetTodo(responseWriter http.ResponseWriter, request *http.Request) error {
	id := chi.URLParam(request, "id")
	todo := getTodo(id)

	err := utils.WriteJSON(responseWriter, http.StatusOK, todo)
	if err != nil {
		log.Fatalf("Failed to write response: %v", err)
	} else {
		fmt.Println("Request completed: GET /todo")
	}

	return nil
}

func (handler TodoHandler) CreateTodo(responseWriter http.ResponseWriter, request *http.Request) error {

	var todo models.TodoCreate
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		log.Fatalf("Error while decoding TODO json")
		return err
	}

	createdTodo := createTodo(todo)
	err = utils.WriteJSON(responseWriter, http.StatusCreated, createdTodo)
	if err != nil {
		log.Fatalf("Failed to write response: %v", err)
	} else {
		fmt.Println("Request completed: POST /todo")
	}

	return nil
}

func (handler TodoHandler) EditTodo(responseWriter http.ResponseWriter, request *http.Request) error {

	var todo models.Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		log.Fatalf("Error while decoding TODO json")
		return err
	}

	editedTodos := editTodo(todo)
	err = utils.WriteJSON(responseWriter, http.StatusCreated, editedTodos)
	if err != nil {
		log.Fatalf("Failed to write response: %v", err)
	} else {
		fmt.Println("Request completed: POST /todo")
	}

	return nil
}

func (handler TodoHandler) DeleteTodo(responseWriter http.ResponseWriter, request *http.Request) error {

	//var id string
	//err := json.NewDecoder(request.Body).Decode(&id)
	id := chi.URLParam(request, "id")
	deletedTodos := deleteTodo(id)

	err := utils.WriteJSON(responseWriter, http.StatusCreated, deletedTodos)
	if err != nil {
		log.Fatalf("Failed to write response: %v", err)
	} else {
		fmt.Println("Request completed: DELETE /todo")
	}

	return nil
}

// Database functions
// --------------------------------------------------------------------------------------------------------------------
func getAllTodos() []models.Todo {

	var todos []models.Todo

	db := storage.NewPostgresStore()
	defer db.Close()

	query := `select * from todo`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error on GET TODO query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo

		// Iterate and makes copies - I think...
		err = rows.Scan(&todo.ID, &todo.Label, &todo.Checked)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		todos = append(todos, todo)
	}

	return todos
}

func getTodo(id string) models.Todo {

	var todo models.Todo

	db := storage.NewPostgresStore()
	defer db.Close()

	query := `select * from todo where id=$1`

	row := db.QueryRow(query, id)

	err := row.Scan(&todo.ID, &todo.Label, &todo.Checked)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No element found with matching ID")
		return todo
	case err == nil:
		return todo
	default:
		log.Fatalf("Unable to scan row: %v", err)
	}

	return todo
}

func createTodo(todo models.TodoCreate) models.Todo {
	db := storage.NewPostgresStore()
	defer db.Close()

	query := `insert into todo (label, checked) values ($1, $2) returning ID, label, checked`

	var returnedTodo models.Todo

	err := db.
		QueryRow(query, todo.Label, todo.Checked).
		Scan(&returnedTodo.ID, &returnedTodo.Label, &returnedTodo.Checked)

	if err != nil {
		log.Fatalf("Error while executing query: %v", err)
	}

	return returnedTodo
}

func editTodo(todo models.Todo) int64 {
	db := storage.NewPostgresStore()
	defer db.Close()

	query := `update todo
		set label=$1, checked=$2
		where id=$3 
	`

	res, err := db.Exec(query, todo.Label, todo.Checked, todo.ID)
	if err != nil {
		log.Fatalf("Error while executing query: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while scanning rows: %v", err)
	}

	return rowsAffected
}

func deleteTodo(id string) int64 {
	db := storage.NewPostgresStore()
	defer db.Close()

	query := `delete from todo where id=$1`

	res, err := db.Exec(query, id)
	if err != nil {
		log.Fatalf("Error while executing query: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while scanning rows: %v", err)
	}

	return rowsAffected
}
