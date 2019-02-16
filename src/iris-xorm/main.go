package main

import (
	"iris-xorm/bootstrap"
	"iris-xorm/web/middleware/identity"
	"iris-xorm/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstar database", "Yourtion")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
