package actions

import "github.com/gobuffalo/buffalo"

// ItemsList default implementation.
func ItemsList(c buffalo.Context) error {
	return c.Render(200, r.HTML("items/list.html"))
}

// ItemsIndex default implementation.
func ItemsIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("items/index.html"))
}
