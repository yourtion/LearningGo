package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"time"
)

var userList []string

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()
	userList = []string{}

	_ = app.Run(iris.Addr("127.0.0.1:8080"))
}

func (c *lotteryController) Get() string {
	count := len(userList)
	return fmt.Sprintf("当前总共参与抽奖用户数： %d\n", count)
}

// POST http://127.0.0.1:8080/import
// params: users
func (c *lotteryController) PostImport() string {
	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",")

	countOrg := len(userList)
	for _, u := range users {
		u := strings.TrimSpace(u)
		if len(u) > 0 {
			userList = append(userList, u)
		}
	}
	countAfter := len(userList)

	return fmt.Sprintf("当前总共参与抽奖用户数： %d，成功导入用户数：%d\n", countAfter, countAfter-countOrg)
}

func (c *lotteryController) GetLuck() string {
	count := len(userList)
	if count > 1 {
		seed := time.Now().UnixNano()
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count))
		user := userList[index]
		userList = append(userList[0:index], userList[index+1:]...)
		return fmt.Sprintf("当前中奖用户：%s，剩余用户数：%d\n", user, count-1)
	} else if count == 1 {
		user := userList[0]
		userList = []string{}
		return fmt.Sprintf("当前中奖用户：%s，剩余用户数：%d\n", user, count-1)
	}
	return fmt.Sprintf("已经没有参与用户了，请先通过 /import 导入\n")
}
