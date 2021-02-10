package main

import (
	"fmt"
	"github.com/kataras/iris/httptest"
	"sync"
	"testing"
)

func TestMVC(t *testing.T) {
	e := httptest.New(t, newApp())

	var wg sync.WaitGroup

	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖用户数： 0\n")

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			e.POST("/import").
				WithFormField("users", fmt.Sprintf("test_u%d", i)).
				Expect().Status(httptest.StatusOK)
		}(i)
	}

	wg.Wait()

	// 验证用户添加成功
	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖用户数： 100\n")
	// 执行抽奖
	e.GET("/luck").Expect().Status(httptest.StatusOK)
	// 验证抽奖后用户数量
	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖用户数： 99\n")
}
