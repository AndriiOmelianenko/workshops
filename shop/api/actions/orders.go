package actions

import "github.com/gobuffalo/buffalo"

// OrdersCreate default implementation.
func OrdersCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("orders/create.html"))
}

// OrdersUpdate default implementation.
func OrdersUpdate(c buffalo.Context) error {
	return c.Render(200, r.HTML("orders/update.html"))
}
