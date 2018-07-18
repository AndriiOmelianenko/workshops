package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/AndriiOmelianenko/workshops/shop/api/models"
	"fmt"
)

// CategoriesList default implementation.
// curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/categories/list
func CategoriesList(c buffalo.Context) error {
	categories := models.Categories{}
	err := models.DB.All(&categories)
	if err != nil {
		fmt.Println("error getting list of categories:", err)
		return c.Render(404, r.String("404 not found\n%v", err))
	}
	return c.Render(200, r.JSON(&categories))
}

// CategoriesIndex default implementation.
// curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/categories/<category ID>
func CategoriesIndex(c buffalo.Context) error {
	category := models.Category{}
	err := models.DB.Find(&category, c.Param("category"))
	if err != nil {
		fmt.Println("error getting specific category:", err)
		return c.Render(404, r.String("404 not found\n%v", err))
	}
	return c.Render(200, r.JSON(&category))
}
