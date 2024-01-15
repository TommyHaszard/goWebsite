package logic

import (
	"fmt"

	"github.com/TommyHaszard/goWebsite/model"
	"github.com/TommyHaszard/goWebsite/server"
	"github.com/patrickmn/go-cache"
)

func CreateUser(c *server.Context, userModel model.User) error {
	c.Cache.Set("currUser", userModel.Name, cache.NoExpiration)
	userModel.Todos = make(map[string]model.Todo)
	err := c.Cache.Add(userModel.Name, userModel, cache.NoExpiration)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetCurrUser(c *server.Context) (interface{}, error) {
	currUserString, ok := c.Cache.Get("currUser")
	if !ok {
		return nil, fmt.Errorf("no curr user flag set in Cache")
	}
	userCache, ok := c.Cache.Get(currUserString.(string))
	if !ok {
		err := fmt.Errorf("cannot retrieve User: %s from Cache", currUserString.(string))
		return userCache, err
	}
	user, ok := userCache.(model.User)
	if !ok {
		err := fmt.Errorf("passed value %s not of User type cannot retrieve from Cache", currUserString)
		return user, err
	}
	return user, nil
}
