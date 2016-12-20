package controllers

import (
	"github.com/revel/revel"
	"github.com/zuiurs/revel-sample-shop/app/models"
	"github.com/zuiurs/revel-sample-shop/app/routes"
	"gopkg.in/validator.v2"
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

func (c Merchandises) Create() revel.Result {
	merchandise := models.Merchandise{}

	if err := c.BindParams(&merchandise); err != nil {
		return c.HandleBadRequestError(err.Error())
	}

	if err := validator.Validate(&merchandise); err != nil {
		return c.HandleBadRequestError(err.Error())
	}

	if err := DB.Create(&merchandise).Error; err != nil {
		return c.HandleInternalServerError("Record Create Failure")
	}

	r := Response{
		Results: merchandise,
	}
	return c.RenderJson(r)
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
