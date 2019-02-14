package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-xorm/services"
)

type IndexController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *IndexController) Get() mvc.Result {
	// TODO:
	return nil
}

func (c *IndexController) GetById(id int) mvc.Result {
	// TODO:
	return nil
}

func (c *IndexController) GetSearch() mvc.Result {
	// TODO:
	return nil
}
