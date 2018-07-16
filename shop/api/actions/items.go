package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/AndriiOmelianenko/workshops/shop/api/models"
	"fmt"
)

// ItemsList default implementation.
func ItemsList(c buffalo.Context) error {
	items := models.Items{}
	err := models.DB.All(&items)
	if err != nil {
		fmt.Println("error getting list of items:", err)
		return c.Render(404, r.String("404 not found\n%v", err))
	}
	return c.Render(200, r.JSON(&items))
}

// ItemsIndex default implementation.
func ItemsIndex(c buffalo.Context) error {
	item := models.Item{}
	err := models.DB.Find(&item, c.Param("item"))
	if err != nil {
		fmt.Println("error getting specific item:", err)
		return c.Render(404, r.String("404 not found\n%v", err))
	}
	return c.Render(200, r.JSON(&item))
}
