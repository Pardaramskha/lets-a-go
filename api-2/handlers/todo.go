package handlers

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/pardaramskha/lets-a-go/api-2/models"
	"github.com/pardaramskha/lets-a-go/api-2/storage"
	"github.com/pardaramskha/lets-a-go/api-2/utils"
)

type TodoHandler struct{}

func (handler TodoHandler) GetTodos(responseWriter http.ResponseWriter, request *http.Request) error {

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

	err = utils.WriteJSON(responseWriter, http.StatusOK, todos)
	if err != nil {
		log.Fatalf("Failed to write response: %v", err)
	} else {
		fmt.Println("Request completed: GET /todos")
	}

	return nil
}
