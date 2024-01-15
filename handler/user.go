package handler

import (
	"github.com/TommyHaszard/goWebsite/logic"
	"github.com/TommyHaszard/goWebsite/model"
	"github.com/TommyHaszard/goWebsite/server"
	"github.com/TommyHaszard/goWebsite/view/user"
)

// func HandleUserShow(c *server.Context) error {
// 	recievedUser := c.FormValue("userFetch")
// 	slog.Info(recievedUser)
// 	userModel, err := logic.GetCurrUser(c, recievedUser)
// 	if err != nil {
// 		return err
// 	}
// 	return render(c, user.ShowUser(userModel.(model.User)))
// }

func HandleCreateUser(c *server.Context) error {
	userName := c.FormValue("userCreate")
	userModel := model.User{
		Name: userName,
	}
	logic.CreateUser(c, userModel)
	return render(c, user.TodoButton())
}

func HandleUserFetchForm(c *server.Context) error {
	return render(c, user.FetchUserForm())
}

// how we insert middleware, we pass in the handler and then call the handler with the context returning a new handler in the httprouter.handle func
func WithAuth(handler server.HandlerFunc) server.HandlerFunc {
	return func(c *server.Context) error {
		return handler(c)
	}
}
