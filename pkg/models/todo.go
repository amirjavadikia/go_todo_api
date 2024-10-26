package models

import (
	"github.com/amirjavadikia/todo_app_with_react_and_go/pkg/config"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func init() {
	var err error

	// connect to database
	err = config.Connect()

	// handle error of database connection
	if err != nil {
		log.Printf("Could not connect to database: %v", err)
	}

	// get a instance of database
	db = config.GetDb()

	// migrate the struct table to database
	err = db.AutoMigrate(&Todo{}).Error

	// handle error of migration
	if err != nil {
		log.Printf("Could not migrate database: %v", err)
	}

	// log success if connection and migration are connect successfully
	log.Println("Database connected and migrated successfully!")
}

func (t *Todo) CreateTodo() (*Todo, error) {

	// add data to the database
	result := db.Create(t)

	// handle error while adding data to database
	if result.Error != nil {
		log.Printf("Error creating todo: %v", result.Error)
		return nil, result.Error
	}

	// return success message and todos when it created successfully
	log.Printf("Todo created successfuly: %v", t)
	return t, nil
}

func GetAll() ([]Todo, error) {
	var todo []Todo
	result := db.Find(&todo)
	return todo, result.Error
}

func GetTodoById(id int64) (*Todo, error) {
	var todo Todo
	result := db.Where("ID = ?", id).Find(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func DeleteTodoById(id int64) error {
	result := db.Delete(&Todo{}, id)
	return result.Error
}
