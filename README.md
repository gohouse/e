
[![GoDoc](https://godoc.org/github.com/gohouse/e?status.svg)](https://godoc.org/github.com/gohouse/e)
[![Go Report Card](https://goreportcard.com/badge/github.com/gohouse/e)](https://goreportcard.com/report/github.com/gohouse/e)
[![GitHub release](https://img.shields.io/github/release/gohouse/e.svg)](https://github.com/gohouse/e/releases/latest)

# e

golang自定义错误包, 可以记录调用位置的具体信息, 包括文件名, 行号, 方法名和错误消息  
cutome error in golang

## 安装

- go mod 
```shell script
require github.com/gohouse/e master
```

- go get
```shell
go get github.com/gohouse/e
```

## 使用

### 生成一个错误
- 标准生成返回
```go
err := e.New("这是错误信息")
// 可选第二个参数,设置记录堆栈层数,默认1层
//err := e.New("这是错误信息", 2)
```

- 附带接受原始错误的返回
```go
err := e.NewWithError("这是错误信息", error.New("这是原生错误信息"))
```

### 获取原生标准错误信息
```go
err.Error()
```

### 获取错误信息包括堆栈
```go
err.ErrorWithStack()
```

### 获取错误堆栈对象`e.ErrorStack`
```go
// 获取堆栈对象
errorStack := err.Stack()
```

### 完整示例
```go
package main

import (
	"fmt"
	"github.com/gohouse/e"
)

func main() {
	var err e.Error
	// 或者 var err e.Error
	err = testError()

	fmt.Println("error msg:", err.Error())
	fmt.Println("error stack:", err.Stack())
	fmt.Println("error with stack:", err.ErrorWithStack())

	fmt.Printf("%#v", err)
}

func testError() e.Error {
	return e.New("only show a custom errors demo", 3)
}
```
输出:  
error msg
```bash
error msg: only show a custom errors demo
```
error stack
```shell script
error stack: [{/go/src/github.com/gohouse/e/examples/demo.go 21 main.testError}]
```
error with stack
```shell script
error msg: only show a custom errors demo
main.test1112
    /go/src/github.com/gohouse/e/examples/test.go:20
main.test111
    /go/src/github.com/gohouse/e/examples/test.go:16
main.main
    /go/src/github.com/gohouse/e/examples/test.go:10
```

> 说明: 为了节约内存占用,默认是记录1层堆栈的信息,如果想要更多堆栈层数,只需要在第二个参数设置对应数量即可,如:  
```shell script
func testError() e.Error {
	return e.New("错误了啦 xxx", 3)
}
```
`err`会记录3层堆栈信息,对应的 stack 为:
```shell script
error stack: [{/go/src/github.com/gohouse/e/examples/demo.go 21 main.testError} {/go/src/github.com/gohouse/e/examples/demo.go 11 main.main} {/usr/local/go/src/runtime/proc.go 203 runtime.main}]
```

## 中间件
我们可以对日志添加中间件,比如做持久化处理,或者打印控制台等
```go
// 使用自带的 log 中间件,并设置默认获取错误堆栈层数为3
e.NewErrorContext().Use(e.Log("errors.log")).Setlayer(3)
e.New("3层错误堆栈测试,并持久化到 errors.log 文件")
```
`e.Log()`中间件代码如下
```go

func Log(fileNames ...string) HandlerFunc {
	return func(ctx *ErrorContext) {
		var fileName = "./error-statck.log"
		if len(fileNames) > 0 {
			fileName = fileNames[0]
		}
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		}
		//write to the file
		_, err = fmt.Fprint(f, time.Now().Format("[2006-01-02 15:04:05] "), ctx.ErrorWithStack(), "--------------------------------------------------\n")
		if err != nil {
			panic(err.Error())
		}
		f.Close()
	}
}
```
日志文件记录如下
```shell script
[2019-12-07 12:25:30] 3层错误堆栈测试,并持久化到 errors.log 文件
demopro/model.IfCheckIn
    /go/src/demopro/model/check_in.go:21
demopro/api.CheckIn
    /go/src/demopro/api/check_in.go:40
github.com/gin-gonic/gin.(*Context).Next
    /go/pkg/mod/github.com/gin-gonic/gin@v1.5.0/context.go:147
--------------------------------------------------
[2020-02-17 15:42:38] 错误信息实验
main.test1112
    /go/src/github.com/gohouse/e/examples/test.go:20
main.test111
    /go/src/github.com/gohouse/e/examples/test.go:16
main.main
    /go/src/github.com/gohouse/e/examples/test.go:11
--------------------------------------------------
```
## 自定义中间件
如果我们只是想打印出来,则可以自己定义一个中间件,如下
```go
func ErrorLog() HandlerFunc {
	return func(ctx *ErrorContext) {
		log.Println(ctx.ErrorWithStack())
	}
}
```
使用
```go
e.NewErrorContext().Use(ErrorLog()).Setlayer(3)
e.New("自定义中间件测试")
```
也可以使用任意多个中间件
```go
e.NewErrorContext().Use(xxx(),xxx()).Use(xxx())
```