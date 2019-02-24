/**
 * 支付宝五福
 * 五福的概率来自识别后的参数(AI图片识别MaBaBa)
 * 基础功能：
 * /lucky 只有一个抽奖的接口，奖品信息都是预先配置好的
 * 测试方法：
 * curl "http://localhost:8080/?rate=4,3,2,1,0"
 * curl "http://localhost:8080/lucky?uid=1&rate=4,3,2,1,0"
 * 压力测试：（这里不存在线程安全问题）
 * wrk -t10 -c10 -d 10 "http://localhost:8080/lucky?uid=1&rate=4,3,2,1,0"
 */

package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var logger *log.Logger

type gift struct {
	id      int    // 奖品ID
	name    string // 奖品名称
	pic     string // 奖品图片
	link    string // 奖品链接
	inUse   bool   // 是否使用中
	rate    int    // 中奖概率，万分之N，0-9999
	rateMin int    // 大于等于中奖编码
	rateMax int    // 小于中奖编码
}

const rateMax = 10

func initLogger() {
	f, _ := os.Create(os.TempDir() + "/lottery.log")
	fmt.Printf("%s\n", os.TempDir())
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

// 初始化奖品列表信息（管理后台来维护）
func newGift() *[5]gift {
	giftlist := new([5]gift)
	// 1 实物大奖
	g1 := gift{
		id:      1,
		name:    "富强福",
		pic:     "富强福.jpg",
		link:    "",
		inUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[0] = g1
	// 2 实物小奖
	g2 := gift{
		id:      2,
		name:    "和谐福",
		pic:     "和谐福.jpg",
		link:    "",
		inUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[1] = g2
	// 3 虚拟券，相同的编码
	g3 := gift{
		id:      3,
		name:    "友善福",
		pic:     "友善福.jpg",
		link:    "",
		inUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[2] = g3
	// 4 虚拟券，不相同的编码
	g4 := gift{
		id:      4,
		name:    "爱国福",
		pic:     "爱国福.jpg",
		link:    "",
		inUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[3] = g4
	// 5 虚拟币
	g5 := gift{
		id:      5,
		name:    "敬业福",
		pic:     "敬业福.jpg",
		link:    "",
		inUse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[4] = g5
	return giftlist
}

func giftRate(rate string) *[5]gift {
	giftList := newGift()
	rates := strings.Split(rate, ",")
	ratesLen := len(rate)

	// 整理奖品数据，把rateMin,rateMax根据rate进行编排
	rateStart := 0
	for i, data := range giftList {
		if !data.inUse {
			continue
		}
		grate := 0
		if i < ratesLen { // 避免数组越界
			grate, _ = strconv.Atoi(rates[i])
		}
		giftList[i].rate = grate
		giftList[i].rateMin = rateStart
		giftList[i].rateMax = rateStart + grate
		if giftList[i].rateMax >= rateMax {
			// 号码达到最大值，分配的范围重头再来
			giftList[i].rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += grate
		}
	}
	fmt.Printf("giftlist=%v\n", giftList)
	return giftList
}

func luckyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(rateMax)
	return code
}

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	initLogger()

	return app
}

func main() {
	app := newApp()

	_ = app.Run(iris.Addr("127.0.0.1:8080"))
}

// GET http://localhost:8080/?rate=4,3,2,1,0
func (c *lotteryController) Get() string {
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	giftList := giftRate(rate)
	return fmt.Sprintf("%v\n", giftList)
}

// GET http://localhost:8080/lucky?uid=1&rate=4,3,2,1,0
func (c *lotteryController) GetLucky() map[string]interface{} {
	uid, _ := c.Ctx.URLParamInt("uid")
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	code := luckyCode()
	ok := false
	result := make(map[string]interface{})
	result["success"] = ok

	giftList := giftRate(rate)
	for _, data := range giftList {
		if !data.inUse {
			continue
		}
		if data.rateMin <= int(code) && data.rateMax >= int(code) {
			// 中奖了，抽奖编码在奖品中奖编码范围内
			ok = true
			sendData := data.pic

			if ok {
				// 中奖后，成功得到奖品（发奖成功）
				// 生成中奖纪录
				saveLuckyData(uid, code, data.id, data.name, data.link, sendData)
				result["success"] = ok
				result["uid"] = uid
				result["id"] = data.id
				result["name"] = data.name
				result["link"] = data.link
				result["data"] = sendData
				break
			}
		}
	}

	return result
}

// 记录用户的获奖记录
func saveLuckyData(uid int, code int32, id int, name, link, sendData string) {
	logger.Printf("lucky, uid=%d, code=%d, gift=%d, name=%s, link=%s, data=%s ", uid, code, id, name, link, sendData)
}
