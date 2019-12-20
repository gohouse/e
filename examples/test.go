package main

import (
	"github.com/gohouse/e"
)

func main() {
	e.NewErrorContext().Use(e.LogFile("xxx.log")).Setlayer(3)
	test111()
}

func test111()  {
	test1112()
}

func test1112()  {
	e.New("错误信息实验")
}