package handler

import (
	"github.com/TommyHaszard/goWebsite/server"
	"github.com/TommyHaszard/goWebsite/view/index"
)

func HandleIndex(c *server.Context) error {
	return render(c, index.Index())
}
