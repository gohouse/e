
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
// 可选第二个参数,设置记录堆栈层数,默认2层
//err := e.New("这是错误信息", 1)
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
error stack: [{/go/src/github.com/gohouse/e/examples/demo.go 21 main.testError} {/go/src/github.com/gohouse/e/examples/demo.go 11 main.main} {/usr/local/go/src/runtime/proc.go 203 runtime.main}]
```
error with stack
```shell script
error with stack: only show a custom errors demo, [/go/src/github.com/gohouse/e/examples/demo.go:21, main.testError], [/go/src/github.com/gohouse/e/examples/demo.go:11, main.main], [/usr/local/go/src/runtime/proc.go:203, runtime.main]
```

> 说明:默认是记录3下层堆栈的信息,如果想要自定义堆栈层数,只需要在第二个参数设置对应数量即可,如:  
```shell script
 err := e.New("错误了啦 xxx", 1)
```
`err`会记录1层堆栈信息,对应的 stack 为:
```shell script
error stack: [{/go/src/github.com/gohouse/e/examples/demo.go 21 main.testError}]
```