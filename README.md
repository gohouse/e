# e
golang自定义错误包, 可以记录调用位置的具体信息, 包括文件名, 行号, 方法名和错误消息  
cutome error in golang

## 安装
```shell
go get github.com/gohouse/e
```

## 使用
```go
package main

import (
	"fmt"
	"github.com/gohouse/e"
)

func main() {
	err := e.New("only show a custom errors demo")

	fmt.Println("msg:",err.Msg)
	fmt.Println("file:",err.Stack.File)
	fmt.Println("line:",err.Stack.Line)
	fmt.Println("func name:",err.Stack.FuncName)
	
	fmt.Printf("%#v",err)
}
```
输出
```bash
msg: only show a custom errors demo
file: /go/src/github.com/demo/e.go
line: 9
func name: main.main

e.Error{Msg:"only show a custom errors demo", Stack:e.Stack{Line:9, FuncName:"main.main", File:"/go/src/github.com/demo/e.go"}}
```

## 转换为原生error
```
var err2 error
err2 = e.New(xxx).ToError()
```

## 获取错误信息
```go
err.Error()
// 或者
err.Msg
```