package routes

import (
	"github.com/kataras/iris/mvc"
	"iris-xorm/bootstrap"
	"iris-xorm/services"
	"iris-xorm/web/controllers"
	"iris-xorm/web/middleware"
)

// GetIndexHandler handles the GET: /
// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	superstarService := services.NewSuperstarService()

	index := mvc.New(b.Party("/"))
	index.Register(superstarService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(superstarService)
	admin.Handle(new(controllers.AdminController))

}
