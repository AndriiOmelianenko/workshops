package grifts

import (
	"github.com/AndriiOmelianenko/workshops/shop/api/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
