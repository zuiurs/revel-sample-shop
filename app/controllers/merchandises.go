package controllers

import (
	"fmt"
	"github.com/revel/revel"
)

type MerchandisesController struct {
	ShopController
}

func (c MerchandisesController) Index() revel.Result {
	r := Response{
		Results: "index",
	}
	return c.RenderJson(r)
}

func (c MerchandisesController) Show(id int) revel.Result {
	r := Response{
		Results: fmt.Sprint("show", id),
	}
	return c.RenderJson(r)
}

func (c MerchandisesController) Create() revel.Result {
	r := Response{
		Results: "create",
	}
	return c.RenderJson(r)
}

func (c MerchandisesController) Delete(id int) revel.Result {
	r := Response{
		Results: fmt.Sprint("delete", id),
	}
	return c.RenderJson(r)
}
