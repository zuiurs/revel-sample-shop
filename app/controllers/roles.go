package controllers

import (
	"github.com/revel/revel"
	"github.com/zuiurs/revel-sample-shop/app/models"
	"github.com/zuiurs/revel-sample-shop/app/routes"
)

type Roles struct {
	Application
}

func (c Roles) Index() revel.Result {
	roles := []models.Role{}

	if err := DB.Order("id desc").Find(&roles).Error; err != nil {
		return c.HandleInternalServerError("Record Find Failure")
	}

	return c.Render(roles)
}

func (c Roles) Show(id int) revel.Result {
	role := models.Role{}

	if err := DB.First(&role, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	r := Response{
		Results: role,
	}
	return c.RenderJson(r)
}

func (c Roles) Create(role models.Role) revel.Result {
	c.Validation.Required(role.Name)

	// roles.validate(c.validaton)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Roles.CreatePage())
	}

	if err := DB.Create(&role).Error; err != nil {
		return c.HandleInternalServerError("Record Create Failure")
	}

	return c.Redirect(routes.Roles.Index())
}

func (c Roles) CreatePage() revel.Result {
	return c.Render()
}

func (c Roles) Delete(id int) revel.Result {
	role := models.Role{}

	if err := DB.First(&role, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	if err := DB.Delete(&role).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}

	return c.Redirect(routes.Roles.Index())
}
