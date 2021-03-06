/**
 * 彩票
 * 1 即刮即得型（已知中奖规则，随机获取号码来匹配是否中奖）
 * 得到随机数： http://localhost:8080/
 *
 * 2 双色球自选型（从已知可选号码中选择每一个位置的号码，等待开奖结果）
 * 开奖号码： http://localhost:8080/prize
 * 规则参考： https://cp.cn/kj/ssq.html?agent=700007
 */

package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"time"
)

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

	_ = app.Run(iris.Addr("127.0.0.1:8080"))
}

// 即开即得型 http://localhost:8080/once
func (c *lotteryController) GetOnce() string {
	var prize string
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(10)
	switch {
	case code == 1:
		prize = "一等奖"
	case code >= 2 && code <= 3:
		prize = "二等奖"
	case code >= 4 && code <= 6:
		prize = "三等奖"
	default:
		return fmt.Sprintf("尾号为1获得一等奖\n"+
			"尾号为2或者3获得二等奖\n"+
			"尾号为456获得三等奖\n"+
			"code=%d\n"+
			"很遗憾，没有获奖", code)
	}
	return fmt.Sprintf("尾号为1获得一等奖\n"+
		"尾号为2或者3获得二等奖\n"+
		"尾号为456获得三等奖\n"+
		"code=%d\n"+
		"恭喜你获得：%s", code, prize)
}

// 双色球自选型 http://localhost:8080/prize
func (c *lotteryController) GetPrize() string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	var prize [7]int

	// 6个红色球，1-33
	for i := 0; i < 6; i++ {
		prize[i] = r.Intn(33) + 1
	}
	// 最后一个蓝色球，1-16
	prize[6] = r.Intn(16) + 1

	return fmt.Sprintf("今日开奖号码是：%v\n", prize)
}
