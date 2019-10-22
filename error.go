package e

import (
	"errors"
	"fmt"
	"runtime"
)

type Error struct {
	*ErrorStack
	msg string
}

func New(msg string) Error {
	var err = new(Error)
	err.msg = msg
	return err.resolveCaller()
}

func NewWithError(msg string, goErr error) Error {
	var err = new(Error)
	if goErr != nil {
		msg = fmt.Sprint(msg, ":", goErr.Error())
	}
	err.msg = msg
	return err.resolveCaller()
}

func (err Error) resolveCaller() Error {
	if funcName, file, line, ok := runtime.Caller(2); ok {
		err.ErrorStack = NewErrorStack()
		err.SetFile(file)
		err.SetLine(line)
		err.SetFuncName(runtime.FuncForPC(funcName).Name())
	}
	return err
}

func (err Error) ToError() error {
	return errors.New(err.msg)
}

func (err Error) ToErrorWithStack() error {
	return errors.New(err.ErrorWithStack())
}

func (err Error) Error() string {
	return err.msg
}

func (err Error) ErrorWithStack() string {
	return fmt.Sprintf("%s; %s:%v:%s",
		err.msg,
		err.GetFile(),
		err.GetLine(),
		err.GetFuncName())
}
func (err Error) Stack() ErrorStack {
	return *err.ErrorStack
}
