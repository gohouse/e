package e

import (
	"errors"
	"runtime"
)

type E interface {
	error
	ToError() error
}
type Stack struct {
	File     string
	Line     int
	FuncName string
}
type Error struct {
	Msg   string
	Stack Stack
}

func New(arg string) Error {
	var err = Error{
		Msg: arg,
	}
	funcName, file, line, ok := runtime.Caller(1)
	if (ok) {
		err.Stack = Stack{
			File:     file,
			Line:     line,
			FuncName: runtime.FuncForPC(funcName).Name(),
		}
	}

	return err
}

func (e Error) ToError() error {
	return errors.New(e.Msg)
}

func (e Error) Error() string {
	return e.Msg
}
