package routes

import (
	"github.com/amirjavadikia/todo_app_with_react_and_go/pkg/controllers"
	"github.com/gorilla/mux"
)

var TodoRoutes = func(router *mux.Router) {

	router.HandleFunc("/todo/", controllers.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todos/", controllers.GetAllTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", controllers.UpdateTodoHandler).Methods("PUT")
	router.HandleFunc("/todo/{id}", controllers.DeleteTodoHandler).Methods("DELETE")
}
