package handler

import (
	"github.com/TommyHaszard/goWebsite/server"

	"github.com/TommyHaszard/goWebsite/view/layout"
)

func HandleLayout(c *server.Context) error {
	return render(c, layout.Base())
}
