package e

import (
	"errors"
	"fmt"
)

func ExampleNew() {
	var err Error
	err = New("这是自定义错误信息")
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
	fmt.Println("错误堆栈信息:", err.Stack())
	fmt.Println("错误堆栈信息-文件名:", err.GetFile())
	fmt.Println("错误堆栈信息-行号:", err.GetLine())
	fmt.Println("错误堆栈信息-方法名:", err.GetFuncName())
}
