package actions

import (
	"github.com/gobuffalo/buffalo"
	"encoding/json"
	"github.com/AndriiOmelianenko/workshops/shop/api/models"
	"strconv"
)

// curl -X POST -H "Content-Type: application/json" http://127.0.0.1:8080/orders/create -d '{"status": "Shipped", "sum": 150}'
// OrdersCreate default implementation.
func OrdersCreate(c buffalo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	order := models.Order{}
	err := decoder.Decode(&order)
	if err != nil {
		return c.Render(404, r.String("ERROR: decoding json: %v", err))
	}
	err = models.DB.Create(&order)
	if err != nil {
		return c.Render(404, r.String("ERROR: creating order record: %v", err))
	}
	return c.Render(200, r.String("Order created: %v", order))
}

// curl -X PUT -H "Content-Type: application/json" http://127.0.0.1:8080/orders/3/item -d '{"item_id": "fa0f13d1-af55-4823-8f85-7e3284316b70", "item_cnt": 5, "item_sum":150}'
// OrdersUpdate default implementation.
func OrdersUpdate(c buffalo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	ordered := models.Ordered{}
	err := decoder.Decode(&ordered)
	if err != nil {
		return c.Render(404, r.String("ERROR: decoding json: %v", err))
	}
	orderID, err := strconv.Atoi(c.Param("order"))
	if err != nil {
		return c.Render(404, r.String("ERROR: converting orderID to int: %v", err))
	}
	ordered.OrderID = orderID
	err = models.DB.Create(&ordered)
	if err != nil {
		return c.Render(404, r.String("ERROR: creating ordered record: %v", err))
	}
	return c.Render(200, r.String("Item added to order: %v", ordered))
}
