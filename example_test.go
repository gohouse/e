package e

import (
	"errors"
	"fmt"
)

func ExampleNew() {
	var err Error
	err = New("这是自定义错误信息", 1)
	fmt.Println(err)
}

func ExampleNewWithError() {
	var err Error
	err = NewWithError(
		"这是自定义错误信息",
		errors.New("这是原始错误信息"))
	fmt.Println(err)
}

func ExampleError_Error() {
	var err Error
	err = New("这是自定义错误信息")
	fmt.Println("错误信息:", err.Error())
}

func ExampleError_ErrorWithStack() {
	var err Error
	err = New("这是自定义错误信息")
	fmt.Println("错误信息,包括堆栈信息:", err.ErrorWithStack())
}

func ExampleError_Stack() {
	var err Error
	err = New("这是自定义错误信息")
	fmt.Println("错误堆栈对象:", err.Stack())
}
