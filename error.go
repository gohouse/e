package e

import (
	"errors"
	"fmt"
	"runtime"
)

// Error 错误主结构体
type Error struct {
	*ErrorStack
	msg string
}

// New 初始化错误
func New(msg string) Error {
	var err = new(Error)
	err.msg = msg
	return err.resolveCaller()
}

// NewWithError 初始化错误,并且可以附加原始错误
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

// ToError 转换为官方原始错误, 只包含错误信息, 不包含错误堆栈信息
func (err Error) ToError() error {
	return errors.New(err.msg)
}

// ToErrorWithStack 跟 func (err Error) ToError() error 类似, 包含自定义错误信息和错误堆栈信息
func (err Error) ToErrorWithStack() error {
	return errors.New(err.ErrorWithStack())
}

// Error 实现了官方的错误, 跟官方用法一致, 直接返回错误信息, 不包含错误堆栈信息
func (err Error) Error() string {
	return err.msg
}

// ErrorWithStack 跟 func (err Error) Error() string 类似, 包含自定义错误信息和错误堆栈信息
func (err Error) ErrorWithStack() string {
	return fmt.Sprintf("%s; %s:%v:%s",
		err.msg,
		err.GetFile(),
		err.GetLine(),
		err.GetFuncName())
}

// Stack 获取错误的具体堆栈信息
func (err Error) Stack() ErrorStack {
	return *err.ErrorStack
}
