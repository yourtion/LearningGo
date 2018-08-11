package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is yourtion@gmail.com

email1 is abc@yourtion.com
email2 is yourtion@qq.com.cn
`

func main() {
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	//fmt.Println(match)
	for _, m := range match {
		fmt.Println(m)
	}
}
