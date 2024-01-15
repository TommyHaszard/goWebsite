package handler

import "github.com/TommyHaszard/goWebsite/server"

func WithCache(handler server.HandlerFunc) server.HandlerFunc {
	return func(c *server.Context) error {

		return handler(c)
	}

}
