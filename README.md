
[![GoDoc](https://godoc.org/github.com/gohouse/e?status.svg)](https://godoc.org/github.com/gohouse/e)
[![Go Report Card](https://goreportcard.com/badge/github.com/gohouse/e)](https://goreportcard.com/report/github.com/gohouse/e)
[![GitHub release](https://img.shields.io/github/release/gohouse/e.svg)](https://github.com/gohouse/e/releases/latest)

# e

golang自定义错误包, 可以记录调用位置的具体信息, 包括文件名, 行号, 方法名和错误消息  
cutome error in golang

## 安装

```shell
go get github.com/gohouse/e
```

## 使用

### 生成一个错误
- 标准生成返回
```go
err := e.New("这是错误信息")
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

### 获取错误堆栈信息
```go
// 获取堆栈对象
errorStack := err.Stack()

// 获取错误的文件
errorStack.GetFile()
// 或者 err.GetFile()

// 获取错误的文件行号
errorStack.GetLine()
// 或者 err.GetLine()

// 获取错误的方法名
errorStack.GetFuncName()
// 或者 err.GetFuncName()
```

### 转换为原生error
```
var err2 error
err2 = err.ToError()
// 这里的err就是e.New(xxx), 即e.Error或e.E
```
这里只包含标准错误信息, 不包含stack信息

### 转换为原生error并附带stack信息
```
var err2 error
err2 = err.ToErrorWithStack()
// 这里的err就是e.New(xxx), 即e.Error或e.E
```
这里包含标准错误信息和stack信息,格式为`error.New("标准错误信息; 错误文件:错误行号:错误方法名")`

### 完整示例
```go
package main

import (
	"fmt"
	"github.com/gohouse/e"
)

func main() {
	var err e.E
	// 或者 var err e.Error
	err = testError()

	fmt.Println("error msg:", err.Error())
	fmt.Println("error stack:", err.Stack())
	fmt.Println("error file:", err.Stack().File)
	fmt.Println("error line:", err.Stack().Line)
	fmt.Println("error func name:", err.Stack().FuncName)

	fmt.Printf("%#v", err)
}

func testError() e.Error {
	return e.New("only show a custom errors demo")
}
```
输出
```bash
error msg: only show a custom errors demo
error stack: {21 main.testError /go/src/github.com/gohouse/demo/e.go}
error file: /go/src/github.com/gohouse/demo/e_demo/e.go
error line: 21
error func name: main.testError

e.Error{Msg:"only show a custom errors demo", Stack:e.Stack{Line:21, FuncName:"main.main", File:"/go/src/github.com/demo/e.go"}}
```