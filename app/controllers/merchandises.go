package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/zuiurs/revel-sample-shop/app/models"
	"github.com/zuiurs/revel-sample-shop/app/routes"
)

type Merchandises struct {
	Application
}

func (c Merchandises) Index() revel.Result {
	merchandises := []models.Merchandise{}

	if err := DB.Order("id desc").Find(&merchandises).Error; err != nil {
		return c.HandleInternalServerError("Record Find Failure")
	}

	return c.Render(merchandises)
}

func (c Merchandises) Show(id int) revel.Result {
	merchandise := models.Merchandise{}

	if err := DB.First(&merchandise, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	r := Response{
		Results: merchandise,
	}
	return c.RenderJson(r)
}

func (c Merchandises) Create(merchandise models.Merchandise) revel.Result {
	if c.Request.Method == "GET" {
		return c.Render()
	}

	if err := DB.Create(&merchandise).Error; err != nil {
		return c.HandleInternalServerError("Record Create Failure")
	}

	return c.Redirect(routes.Merchandises.Index())
}

func (c Merchandises) Delete(id int) revel.Result {
	merchandise := models.Merchandise{}

	if err := DB.First(&merchandise, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	if err := DB.Delete(&merchandise).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}

	return c.Redirect(routes.Merchandises.Index())
}

func (c Merchandises) AddCart(id int) revel.Result {
	c.Session["_cart"] += fmt.Sprintf("&%d#%d", id, 1) // 1はいずれ個数に置き換える
	c.Flash.Success(fmt.Sprintf("Add %d", id))
	return c.Redirect(routes.Merchandises.Index())
}
