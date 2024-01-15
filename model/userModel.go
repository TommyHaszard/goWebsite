package model

type User struct {
	Name  string
	Email string
	Todos map[string]Todo
}
