package e

import (
	"errors"
	"fmt"
	"runtime"
)

type E interface {
	error
	ToError() error
	Stack() Stack
}
type Stack struct {
	File     string
	Line     int
	FuncName string
}
type Error struct {
	msg   string
	stack Stack
}

func New(msg string) Error {
	var err = Error{
		msg: msg,
	}
	funcName, file, line, ok := runtime.Caller(1)
	if (ok) {
		err.stack = Stack{
			File:     file,
			Line:     line,
			FuncName: runtime.FuncForPC(funcName).Name(),
		}
	}

	return err
}

func NewWithError(msg string, goErr error) Error {
	var err = Error{
		msg: fmt.Sprint(msg, ":", goErr.Error()),
	}
	funcName, file, line, ok := runtime.Caller(1)
	if (ok) {
		err.stack = Stack{
			File:     file,
			Line:     line,
			FuncName: runtime.FuncForPC(funcName).Name(),
		}
	}

	return err
}

func (e Error) ToError() error {
	return errors.New(e.msg)
}

func (e Error) Error() string {
	return e.msg
}

func (e Error) Stack() Stack {
	return e.stack
}
