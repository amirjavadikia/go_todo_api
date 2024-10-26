package main

import (
	"github.com/amirjavadikia/todo_app_with_react_and_go/pkg/config"
	"github.com/amirjavadikia/todo_app_with_react_and_go/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const port = ":9000"

func main() {

	r := mux.NewRouter()
	routes.TodoRoutes(r)
	log.Printf("Server is running on port %s", port)

	//cors configuration
	handler := config.ConfigureCORS(r)

	// http serve
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Printf("Faild to start server : %v", err)
	}
}
