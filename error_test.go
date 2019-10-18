package e

import (
	"errors"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("only show a custom errors demo")

	t.Log("msg:",err.Msg)
	t.Log("file:",err.Stack.File)
	t.Log("line:",err.Stack.Line)
	t.Log("func name:",err.Stack.FuncName)
	fmt.Printf("%#v",err)
}

func TestNewWithError(t *testing.T) {
	goErr := errors.New("这是官方标准错误")
	err := NewWithError("手动添加的错误", goErr)

	t.Log("msg:",err.Msg)
	t.Log("file:",err.Stack.File)
	t.Log("line:",err.Stack.Line)
	t.Log("func name:",err.Stack.FuncName)
	fmt.Printf("%#v",err)
}

func TestError_Error(t *testing.T) {
	err := New("only show a custom errors demo")

	t.Log(err.Error())
}


