package main

import (
	"github.com/TommyHaszard/goWebsite/handler"
	"github.com/TommyHaszard/goWebsite/server"
)

func main() {
	server := server.New()
	server.Get("/", handler.HandleIndex)
	server.Post("/createUser", handler.HandleCreateUser)
	server.Get("/showTodos", handler.HandleTodoShow)
	server.Post("/createTodo", handler.HandleTodoCreate)
	server.Get("/userFetchForm", handler.HandleUserFetchForm)
	server.Get("/todoFetchForm", handler.HandleTodoForm)
	server.Start(":3000")
}
