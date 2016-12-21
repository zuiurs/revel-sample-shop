package controllers

import (
	"github.com/revel/revel"
	"github.com/zuiurs/revel-sample-shop/app/models"
	"github.com/zuiurs/revel-sample-shop/app/routes"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Application
}

func (c Users) Index() revel.Result {
	users := []models.User{}

	if err := DB.Order("id desc").Find(&users).Error; err != nil {
		return c.HandleInternalServerError("Record Find Failure")
	}

	return c.Render(users)
}

func (c Users) Show(id int) revel.Result {
	user := models.User{}

	if err := DB.First(&user, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	r := Response{
		Results: user,
	}
	return c.RenderJson(r)
}

func (c Users) Create(user models.User) revel.Result {
	c.Validation.Required(user.Username)
	c.Validation.Required(user.Password)

	// roles.validate(c.validaton)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Users.CreatePage())
	}

	user.HashedPassword, _ = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err := DB.Create(&user).Error; err != nil {
		return c.HandleInternalServerError("Record Create Failure")
	}

	return c.Redirect(routes.Users.Index())
}

func (c Users) CreatePage() revel.Result {
	return c.Render()
}

func (c Users) Delete(id int) revel.Result {
	user := models.User{}

	if err := DB.First(&user, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	if err := DB.Delete(&user).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}

	return c.Redirect(routes.Users.Index())
}

func (c Users) Settings(id int) revel.Result {
	return nil
}

func (c Users) SettingsPage(id int) revel.Result {
	return c.Render()
}

func (c Users) ModifyPage(id int) revel.Result {
	roles := []models.Role{}
	if err := DB.Find(&roles).Error; err != nil {
		c.HandleInternalServerError("Record Find Failure")
	}

	user := models.User{}
	if err := DB.First(&user, id).Error; err != nil {
		c.HandleInternalServerError("Record Find Failure")
	}

	return c.Render(roles, user)
}

// 引数はviewで宣言したやつと同じじゃなきゃダメ
func (c Users) Modify(id int, user models.User) revel.Result {
	u := models.User{}
	if err := DB.First(&u, id).Error; err != nil {
		c.HandleInternalServerError("Record Find Failure")
	}

	u.Username = user.Username
	u.HashedPassword, _ = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	u.RoleID = user.RoleID

	if err := DB.Save(&u).Error; err != nil {
		c.HandleInternalServerError("Record Update Failure")
	}

	return c.Redirect(routes.Users.Index())
}
