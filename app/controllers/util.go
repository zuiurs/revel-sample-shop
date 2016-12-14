package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"io/ioutil"
	"net/http"
)

type ShopController struct {
	*revel.Controller
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Results interface{} `json:"results"`
}

func (c *ShopController) BindParams(s interface{}) error {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, s)
}

func (c *ShopController) HandleBadRequestError(s string) revel.Result {
	c.Response.Status = http.StatusBadRequest // 400
	r := ErrorResponse{
		Code:    c.Response.Status,
		Message: s,
	}
	return c.RenderJson(r)
}

func (c *ShopController) HandleNotFoundError(s string) revel.Result {
	c.Response.Status = http.StatusNotFound // 404
	r := ErrorResponse{
		Code:    c.Response.Status,
		Message: s,
	}
	return c.RenderJson(r)
}

func (c *ShopController) HandleInternalServerError(s string) revel.Result {
	c.Response.Status = http.StatusInternalServerError // 500
	r := ErrorResponse{
		Code:    c.Response.Status,
		Message: s,
	}
	return c.RenderJson(r)
}
