/**
 * 微博抢红包
 * 两个步骤：
 * 1. 抢红包，设置红包总金额，红包个数，返回抢红包的地址
 * curl "http://localhost:8080/set?uid=1&money=100&num=100"
 * 2. 抢红包，先到先得，随机得到红包金额
 * curl "http://localhost:8080/get?id=1&uid=1"
 * 压力测试：
 * wrk -t10 -c10 -d5  "http://localhost:8080/set?uid=1&money=100&num=10"
 */
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"sync"
	"time"
)

type task struct {
	id       uint32
	callback chan uint
}

//var packageList = make(map[uint32][]uint)
var packageList = new(sync.Map)

//var chTasks = make(chan task)
const taskNum = 16

var chTasksList = make([]chan task, taskNum)

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	for i := 0; i < taskNum; i++ {
		chTasksList[i] = make(chan task)
		go fetchPackageListMoney(chTasksList[i])

	}

	return app
}

func main() {
	app := newApp()

	_ = app.Run(iris.Addr("127.0.0.1:8080"))
}

// GET http://localhost:8080/
func (c *lotteryController) Get() map[uint32][2]int {
	result := make(map[uint32][2]int)
	packageList.Range(func(key, value interface{}) bool {
		id := key.(uint32)
		list := value.([]uint)
		var money int
		for _, v := range list {
			money += int(v)
		}
		result[id] = [2]int{len(list), money}
		return true
	})
	return result
}

// 发红包
// GET http://localhost:8080/set?uid=1&money=100&num=100
func (c *lotteryController) GetSet() string {
	uid, errUid := c.Ctx.URLParamInt("uid")
	money, errMoney := c.Ctx.URLParamFloat64("money")
	num, errNum := c.Ctx.URLParamInt("num")
	if errUid != nil || errMoney != nil || errNum != nil {
		return fmt.Sprintf("参数格式异常，errUid=%s, errMoney=%s, errNum=%s\n", errUid, errMoney, errNum)
	}

	moneyTotal := int(money * 100)
	if uid < 1 || moneyTotal < num || num < 1 {
		return fmt.Sprintf("参数数值异常，uid=%d, money=%f, num=%d\n", uid, money, num)
	}

	// 分配红包金额
	seed := time.Now().UnixNano()
	rMax := 0.55 // 随机分配的最大值
	// 根据红包数量调整最大红包金额
	if num > 1000 {
		rMax = 0.01
	} else if num >= 100 {
		rMax = 0.1
	} else if num >= 10 {
		rMax = 0.3
	}
	list := make([]uint, num)
	leftMoney := moneyTotal
	leftNum := num
	r := rand.New(rand.NewSource(seed))

	// 大循环开始，只要还有没分配的名额，继续分配
	for leftNum > 0 {

		if leftNum == 1 {
			// 最后一个名额，把剩余的全部给它
			list[num-1] = uint(leftMoney)
			break
		}

		// 剩下的最多只能分配到1分钱时，不用再随机
		if leftMoney == leftNum {
			for i := num - leftNum; i < num; i++ {
				list[i] = 1
			}
			break
		}

		// 每次对剩余金额的1%-55%随机，最小1，最大就是剩余金额55%（需要给剩余的名额留下1分钱的生存空间）
		rMoney := int(float64(leftMoney-leftNum) * rMax)
		m := r.Intn(rMoney)
		if m < 1 {
			m = 1
		}
		list[num-leftNum] = uint(m)
		leftMoney -= m
		leftNum--
	}

	// 红包的唯一ID
	id := r.Uint32()
	packageList.Store(id, list)
	// 返回抢红包的URL
	return fmt.Sprintf("/get?id=%d&uid=%d&num=%d\n", id, uid, num)
}

// 抢红包
// GET http://localhost:8080/get?id=1&uid=1
func (c *lotteryController) GetGet() string {
	uid, errUid := c.Ctx.URLParamInt("uid")
	id, errId := c.Ctx.URLParamInt("id")
	if errUid != nil || errId != nil {
		return fmt.Sprintf("参数格式异常，errUid=%s, errId=%s\n", errUid, errId)
	}
	if uid < 1 || id < 1 {
		return fmt.Sprintf("参数数值异常，uid=%d, id=%d\n", uid, id)
	}

	listLoad, ok := packageList.Load(uint32(id))
	if !ok || listLoad == nil {
		return fmt.Sprintf("红包不存在,id=%d\n", id)
	}
	list := listLoad.([]uint)
	if len(list) < 1 {
		return fmt.Sprintf("红包不存在,id=%d\n", id)
	}

	// 构造一个抢红包任务
	callback := make(chan uint)
	t := task{id: uint32(id), callback: callback}
	// 把任务发送给channel
	chTask := chTasksList[id%taskNum]
	chTask <- t
	// 回调的channel等待处理结果
	money := <-callback
	if money <= 0 {
		fmt.Println(uid, "很遗憾，没能抢到红包")
		return fmt.Sprintf("很遗憾，没能抢到红包\n")
	} else {
		fmt.Println(uid, "抢到一个红包，金额为:", money)
		return fmt.Sprintf("恭喜你抢到一个红包，金额为:%d\n", money)
	}
}

func fetchPackageListMoney(chTask chan task) {
	for {
		t := <-chTask
		id := t.id
		listLoad, ok := packageList.Load(uint32(id))
		if ok && listLoad != nil {
			list := listLoad.([]uint)
			// 分配的随机数
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			// 从红包金额中随机得到一个
			i := r.Intn(len(list))
			money := list[i]
			// 更新红包列表中的信息
			if len(list) > 1 {
				if i == len(list)-1 {
					packageList.Store(uint32(id), list[:i])
				} else if i == 0 {
					packageList.Store(uint32(id), list[1:])
				} else {
					packageList.Store(uint32(id), append(list[:i], list[i+1:]...))
				}
			} else {
				packageList.Delete(uint32(id))
			}
			t.callback <- money
		} else {
			t.callback <- 0
		}
	}
}
