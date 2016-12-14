package controllers

import (
	"github.com/revel/revel"
)

type SystemController struct {
	ShopController
}

func (c SystemController) Login() revel.Result {
	r := Response{
		Results: "login",
	}
	return c.RenderJson(r)
}

func (c SystemController) Logout() revel.Result {
	r := Response{
		Results: "logout",
	}
	return c.RenderJson(r)
}
