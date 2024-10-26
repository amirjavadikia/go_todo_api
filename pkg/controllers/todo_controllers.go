package controllers

import (
	"github.com/amirjavadikia/todo_app_with_react_and_go/pkg/config"
	"github.com/amirjavadikia/todo_app_with_react_and_go/pkg/models"
	"github.com/amirjavadikia/todo_app_with_react_and_go/pkg/utils"
	"github.com/amirjavadikia/todo_app_with_react_and_go/pkg/validators"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	// get an instance of todo model
	var todo models.Todo
	// convert request body to json and error handling
	if err := utils.ParseJSON(w, r, &todo); err != nil {
		return
	}

	if err := validators.ValidateTodoInput(todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Create a new record to database
	createTodo, err := todo.CreateTodo()
	// Error handling while new record add to database
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Send success response and data
	utils.SendResponse(w, http.StatusCreated, createTodo)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendResponse(w, http.StatusOK, todos)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var updatedData models.Todo
	if err := utils.ParseJSON(w, r, &updatedData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	id := params["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	todoDetails, err := models.GetTodoById(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if updatedData.Title != "" {
		todoDetails.Title = updatedData.Title
	}

	if saveErr := config.GetDb().Save(&todoDetails).Error; saveErr != nil {
		http.Error(w, saveErr.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, http.StatusOK, todoDetails)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ID, _ := strconv.ParseInt(id, 0, 0)
	err := models.DeleteTodoById(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendResponse(w, http.StatusOK, "Todo deleted successfully")
}
