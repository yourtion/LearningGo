package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-xorm/services"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *AdminController) Get() mvc.Result {
	// TODO:
	return nil
}

func (c *AdminController) GetEdit() mvc.Result {
	// TODO:
	return nil
}

func (c *AdminController) PostSave() mvc.Result {
	// TODO:
	return nil
}

func (c *AdminController) GetDelete() mvc.Result {
	// TODO:
	return nil
}
