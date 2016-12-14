package controllers

import (
	"fmt"
	"github.com/revel/revel"
)

type RolesController struct {
	ShopController
}

func (c RolesController) Index() revel.Result {
	r := Response{
		Results: "index",
	}
	return c.RenderJson(r)
}

func (c RolesController) Show(id int) revel.Result {
	r := Response{
		Results: fmt.Sprint("show", id),
	}
	return c.RenderJson(r)
}

func (c RolesController) Create() revel.Result {
	r := Response{
		Results: "create",
	}
	return c.RenderJson(r)
}

func (c RolesController) Delete(id int) revel.Result {
	r := Response{
		Results: fmt.Sprint("delete", id),
	}
	return c.RenderJson(r)
}
