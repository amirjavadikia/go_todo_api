package validators

import (
	"github.com/amirjavadikia/todo_app_with_react_and_go/pkg/models"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type TodoValidate struct {
	Title     string `validate:"required,min=3,max=100"`
	Completed bool   `validate:"required"`
}

func ValidateTodoInput(todo models.Todo) error {
	todoValidator := TodoValidate{
		Title:     todo.Title,
		Completed: todo.Completed,
	}
	return validate.Struct(todoValidator)
}
