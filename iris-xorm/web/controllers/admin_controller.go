package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-xorm/models"
	"iris-xorm/services"
	"log"
	"time"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *AdminController) Get() mvc.Result {
	dataList := c.Service.GetAll()
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"DataList": dataList,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminController) GetEdit() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	var data *models.StarInfo
	if err == nil {
		data = c.Service.Get(id)
	} else {
		data = &models.StarInfo{Id: 0}
	}
	return mvc.View{Name: "admin/edit.html",
		Data: iris.Map{
			"Title": "管理后台",
			"info":  data,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminController) PostSave() mvc.Result {
	info := models.StarInfo{}
	err := c.Ctx.ReadForm(&info)
	if err != nil {
		log.Fatal(err)
	}
	if info.Id > 0 {
		info.SysUpdated = int(time.Now().Unix())
		err = c.Service.Update(
			&info,
			[]string{"name_zh", "name_en", "avatar", "birthday", "height", "weight", "club", "jersey", "country", "more_info"},
		)
	} else {
		info.SysCreated = int(time.Now().Unix())
		err = c.Service.Create(&info)
	}
	return mvc.Response{Path: "/admin"}
}

func (c *AdminController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		err = c.Service.Delete(id)
	}
	return mvc.Response{Path: "/admin"}
}
