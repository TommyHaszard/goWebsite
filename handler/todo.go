package handler

import (
	"github.com/TommyHaszard/goWebsite/logic"
	"github.com/TommyHaszard/goWebsite/model"
	"github.com/TommyHaszard/goWebsite/server"
	"github.com/TommyHaszard/goWebsite/view/todo"
)

func HandleTodoCreate(c *server.Context) error {
	name := c.FormValue("name")
	desc := c.FormValue("desc")
	err := logic.CreateTodo(c, name, desc)
	if err != nil {
		return err
	}
	return HandleTodoShow(c)
}

func HandleTodoDelete(c *server.Context) error {
	return nil
}

func HandleTodoEdit(c *server.Context) error {
	return nil
}

func HandleTodoShow(c *server.Context) error {
	inter, err := logic.GetCurrUser(c)
	if err != nil {
		return err
	}
	user := inter.(model.User)
	return render(c, todo.ShowTodos(user.Todos))
}

func HandleTodoForm(c *server.Context) error {
	return render(c, todo.CreateTodo())
}
