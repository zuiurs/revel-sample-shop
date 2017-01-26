package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/zuiurs/revel-sample-shop/app/models"
	"github.com/zuiurs/revel-sample-shop/app/routes"
)

type Orders struct {
	Application
}

func (c Orders) Index() revel.Result {
	orders := []models.Order{}

	if err := DB.Order("id desc").Find(&orders).Error; err != nil {
		return c.HandleInternalServerError("Record Find Failure")
	}

	return c.Render(orders)
}

func (c Orders) Delete(id int) revel.Result {
	order := models.Order{}

	if err := DB.First(&order, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	if err := DB.Delete(&order).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}

	return c.Redirect(routes.Orders.Index())
}

func (c Orders) Complete() revel.Result {
	list, err := models.ParseCart(c.Session["_cart"])
	if err != nil {
		return c.HandleInternalServerError(fmt.Sprintf("Cart Session Parse Error: %s", err))
	}
	c.Session["_cart"] = ""
	fmt.Println(list)

	lo := models.Order{}
	DB.Last(&lo)

	orderID := lo.OrderID + 1

	for k, v := range list {
		o := models.Order{
			OrderID: orderID,
			ItemID:  uint64(k),
			Num:     int64(v),
		}

		if err := DB.Create(&o).Error; err != nil {
			return c.HandleInternalServerError("Record Create Failure")
		}
	}

	return c.Redirect(routes.Orders.Index())
}
