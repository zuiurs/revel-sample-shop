package controllers

import (
	"fmt"
	"github.com/revel/revel"
)

type UsersController struct {
	ShopController
}

func (c UsersController) Index() revel.Result {
	r := Response{
		Results: "index",
	}
	return c.RenderJson(r)
}

func (c UsersController) Show(id int) revel.Result {
	r := Response{
		Results: fmt.Sprint("show", id),
	}
	return c.RenderJson(r)
}

func (c UsersController) Create() revel.Result {
	r := Response{
		Results: "create",
	}
	return c.RenderJson(r)
}

func (c UsersController) Delete(id int) revel.Result {
	r := Response{
		Results: fmt.Sprint("delete", id),
	}
	return c.RenderJson(r)
}
