package main

import (
	"github.com/gohouse/e"
	"log"
)

func main() {
	e.NewErrorContext().Use(e.LogFile("xxx.log")).Setlayer(3)
	err := test111()
	log.Println(err.Stack())
	log.Println(err.ErrorWithStack())
}

func test111() e.Error {
	return test1112()
}

func test1112() e.Error {
	return e.New("错误信息实验")
}