/**
 * 大转盘程序
 * curl http://localhost:8080/
 * curl http://localhost:8080/debug
 * curl http://localhost:8080/prize
 * 固定几个奖品，不同的中奖概率或者总数量限制
 * 每一次转动抽奖，后端计算出这次抽奖的中奖情况，并返回对应的奖品信息
 *
 * 不用互斥锁，而是用CAS操作来更新，保证并发库存更新的正常
 * 压力测试：
 * wrk -t10 -c100 -d5 "http://localhost:8080/prize"
 */

package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"
)

// 奖品中奖概率
type prizeRate struct {
	Rate  int    // 万分之N的中奖概率
	Total int    // 总数量限制，0 表示无限数量
	CodeA int    // 中奖概率起始编码（包含）
	CodeB int    // 中奖概率终止编码（包含）
	Left  *int32 // 剩余数
}

// 奖品列表
var prizeList = []string{
	"一等奖，火星单程船票",
	"二等奖，凉飕飕南极之旅",
	"三等奖，iPhone一部",
	"", // 没有中奖
}

// 奖品的中奖概率设置，与上面的 prizeList 对应的设置
var left = int32(1000)
var rateList = []prizeRate{
	{100, 1000, 0, 9999, &left},
	//{20, 2, 1, 2, 2},
	//{50, 10, 3, 5, 10},
	//{1000, 0, 100, 9999, 1000},
}

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

func (c *lotteryController) Get() string {
	return fmt.Sprintf("大转盘奖品列表：\n%s", strings.Join(prizeList, "\n"))
}

func (c *lotteryController) GetDebug() string {
	return fmt.Sprintf("获奖概率：\n%v \n", rateList)
}

// GET http://localhost:8080/prize
func (c *lotteryController) GetPrize() string {

	// 第一步，根据随机数匹配奖品
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	code := r.Intn(10000)

	//fmt.Println("GetPrize code=", code)
	var myprize string
	var prizeRate *prizeRate
	// 从奖品列表中匹配，是否中奖
	for i, prize := range prizeList {
		rate := &rateList[i]
		if code >= rate.CodeA && code <= rate.CodeB {
			// 满足中奖条件
			myprize = prize
			prizeRate = rate
			break
		}
	}
	if myprize == "" {
		// 没有中奖
		myprize = "很遗憾，再来一次"
		return myprize
	}
	// 第二步，发奖，是否可以发奖
	if prizeRate.Total == 0 {
		// 无限奖品
		fmt.Println("中奖： ", myprize)
		return myprize
	} else if *prizeRate.Left > 0 {
		// 还有剩余奖品，使用 CAS 操作来做安全更新
		left := atomic.AddInt32(prizeRate.Left, -1)
		if left >= 0 {
			fmt.Println("中奖： ", myprize)
			return myprize
		}
	}
	// 有限且没有剩余奖品，无法发奖
	myprize = "很遗憾，再来一次"
	return myprize
}
