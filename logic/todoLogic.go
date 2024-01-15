package logic

import (
	"github.com/TommyHaszard/goWebsite/model"
	"github.com/TommyHaszard/goWebsite/server"
	"github.com/patrickmn/go-cache"
)

func CreateTodo(c *server.Context, name string, description string) error {
	todo := model.Todo{
		Name:        name,
		Description: description,
		Completed:   false,
	}
	inter, err := GetCurrUser(c)
	if err != nil {
		return err
	}

	userModel := inter.(model.User)
	userModel.Todos[todo.Name] = todo

	// add user back to cache with new Todo
	c.Cache.Set(userModel.Name, userModel, cache.NoExpiration)
	return nil
}
