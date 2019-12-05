package main

import (
	"fmt"
	"github.com/gohouse/e"
)

func main() {
	var err e.Error
	// 或者 var err e.Error
	err = test1()

	fmt.Println("error msg:", err.Error())
	fmt.Println("error stack:", err.Stack())
	fmt.Println("error stack:", err.ErrorWithStack())
}

func test1() e.Error {
	return test2()
}

func test2() e.Error {
	return test3()
}

func test3() e.Error {
	return e.New("xxxxxxx", 3)
}
