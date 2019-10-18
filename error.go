package e

import (
	"errors"
	"runtime"
)

type E interface {
	ToError() error
	Error() string
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
		//return errors.New(fmt.Sprintf("msg=%s file=%s line=%v funcName=%s",
		//	arg, file, line, runtime.FuncForPC(funcName).Name()))
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
