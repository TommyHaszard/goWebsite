package handler

import (
	"github.com/TommyHaszard/goWebsite/server"
	"github.com/a-h/templ"
)

func render(c *server.Context, component templ.Component) error {
	return component.Render(c.Request.Context(), c.Response)
}
