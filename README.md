# e
golang自定义错误包, 可以记录调用位置的具体信息, 包括文件名, 行号, 方法名和错误消息  
cutome error in golang

## 安装
```shell
go get github.com/gohouse/e
```

## 使用

### 转换为原生error
```
var err2 error
err2 = e.New(xxx).ToError()
```

### 获取错误信息
```go
err.Error()
```

### 获取错误堆栈信息
```go
errorStack := err.Stack()
errorStack.File
errorStack.Line
errorStack.FuncName
```

### 完整示例
```go
package main

import (
	"fmt"
	"github.com/gohouse/e"
)

func main() {
	err := testError()

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
error file: /Users/fizz/go/src/github.com/gohouse/demo/e_demo/e.go
error line: 21
error func name: main.testError

e.Error{Msg:"only show a custom errors demo", Stack:e.Stack{Line:9, FuncName:"main.main", File:"/go/src/github.com/demo/e.go"}}
```