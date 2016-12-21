package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/zuiurs/revel-sample-shop/app/models"
	"github.com/zuiurs/revel-sample-shop/app/routes"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const (
	DefaultRoleID   = 13
	DefaultRoleName = "staff"

	AdminRoleID   = 12
	AdminRoleName = "admin"
)

type Application struct {
	ShopController
}

// store current user struct to session
func (c Application) HoldCurrentUser() revel.Result {
	if user := c.connected(); user != nil {
		c.RenderArgs["currentUser"] = user
	}
	return nil
}

// handling by user's role_id
func (c Application) CheckUser() revel.Result {
	if c.Action == "Application.Index" {
		return nil
	}

	if user := c.connected(); user != nil {
		// admin 以外のユーザーの下記のアクセスは拒否
		if user.RoleID != AdminRoleID {
			switch c.Action {
			// black list
			case
				"Merchandises.Delete",
				"Merchandises.Create",
				"Users.Index",
				"Users.Show",
				"Users.Delete",
				"Roles.Index",
				"Roles.Show",
				"Roles.Create",
				"Roles.Delete",
				"Application.Admin":
				return c.failedAccessAdminOnlyPage()
			}
		}
	} else {
		// loginable route via guest mode
		switch c.Action {
		case "Merchandises.Index":
			// OK
			c.Flash.Error("Guest mode has some restrictions.")
			c.FlashParams()
		default:
			if strings.HasPrefix(c.Action, "Application") {
				//application.admin 以外の application メソッドだったら通す。
				if !strings.HasPrefix(c.Action, "Application.Admin") {
					break
				} else {
					// NG
					c.Flash.Error("Please Login first")
					return c.Redirect(routes.Application.Index())
				}
			} else {
				// NG
				c.Flash.Error("Please Login first")
				return c.Redirect(routes.Application.Index())
			}
		}
	}
	return nil
}

func (c Application) failedAccessAdminOnlyPage() revel.Result {
	c.Flash.Error("cannot access admin only page")
	return c.Redirect(routes.Application.Index())
}

func (c Application) connected() *models.User {
	// check whether current user set
	if c.RenderArgs["currentUser"] != nil {
		return c.RenderArgs["currentUser"].(*models.User)
	}
	// check session cookie
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username)
	}
	return nil
}

func (c Application) getUser(username string) *models.User {
	user := models.User{}
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func (c Application) Index() revel.Result {
	if c.connected() != nil {
		return c.Redirect(routes.Merchandises.Index())
	}

	c.Flash.Error("Please login first")
	return c.Render()
}

func (c Application) Login(username, password string) revel.Result {
	user := c.getUser(username)

	if user != nil {
		if err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password)); err == nil {
			c.Session["user"] = username
			c.Flash.Success(fmt.Sprintf("Welcome to goyasan, %s!", username))
			return c.Redirect(routes.Merchandises.Index())
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(routes.Application.LoginPage())
}

func (c Application) LoginPage() revel.Result {
	return c.Render()
}

func (c Application) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.Application.Index())
}

func (c Application) Register(user models.User, verifyPassword string) revel.Result {
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == user.Password).Message("Not same password")

	// user.validate(c.validaton)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Application.RegisterPage())
	}

	user.HashedPassword, _ = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err := DB.Create(&user).Error; err != nil {
		c.HandleInternalServerError(err.Error())
	}

	user.RoleID = DefaultRoleID

	c.Session["user"] = user.Username
	c.Flash.Success(fmt.Sprintf("Welcome to goyasan, %s!", user.Username))
	return c.Redirect(routes.Merchandises.Index())
}

func (c Application) RegisterPage() revel.Result {
	return c.Render()
}

func (c Application) Admin() revel.Result {
	//admin role id の検査を行う
	return c.Render()
}

func (c Application) AdminPower() revel.Result {
	if user := c.connected(); user != nil {
		user.RoleID = AdminRoleID
		if err := DB.Save(&user).Error; err != nil {
			c.HandleInternalServerError("failed to update status")
		}
	}
	return c.Redirect(routes.Application.Admin())
}
