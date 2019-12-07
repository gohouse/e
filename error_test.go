package e

import (
	"errors"
	"fmt"
	"testing"
)

func _print(t *testing.T, err Error) {
	t.Log("msg:", err.Error())
	_printStack(t, err.Stack())
	fmt.Printf("%#v\n", err)
}

func _printStack(t *testing.T, stack ErrorStackChain) {
	for _, item := range stack {
		t.Log(fmt.Sprintf("%s,%s:%v", item.FuncName, item.File, item.Line))
	}
}

func TestNew(t *testing.T) {
	err := New("only show a custom errors demo")

	_print(t, err)
}

func TestNewWithError(t *testing.T) {
	goErr := errors.New("这是官方标准错误")
	err := NewWithError("手动添加的错误", goErr)

	_print(t, err)
}

func TestError_Error(t *testing.T) {
	err := New("only show a custom errors demo")

	t.Log(err.Error())
}

func TestError_ErrorWithStack(t *testing.T) {
	goErr := errors.New("这是官方标准错误")
	err := NewWithError("手动添加的错误", goErr)

	t.Log(err.Error())
}

func TestError_Stack(t *testing.T) {
	err := New("only show a custom errors demo")
	_printStack(t, err.Stack())
}
