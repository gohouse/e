// because the former will succeed if err wraps an *os.PathError.
package e

import (
	"fmt"
	"runtime"
)

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string, depth ...int) *errorString {
	var dp = 2
	if len(depth) > 0 {
		dp = depth[0]
	}
	err := &errorString{msg: text, ErrorStack: NewErrorStack(), depth: dp}
	return err.resolveCaller()
}

// errorString is a trivial implementation of error.
type errorString struct {
	ErrorStack
	msg   string
	depth int
}

// NewWithError 初始化错误,并且可以附加原始错误
func NewWithError(msg string, goErr error, depth ...int) *errorString {
	if goErr != nil {
		msg = fmt.Sprint(msg, ":", goErr.Error())
	}
	return New(msg, depth...)
}

func (err *errorString) resolveCaller() *errorString {
	// 取三层
	var i = 0
	for i < err.depth {
		if funcName, file, line, ok := runtime.Caller(i + 2); ok {
			var obj = StackObject{
				File:     file,
				Line:     line,
				FuncName: runtime.FuncForPC(funcName).Name(),
			}
			err.ErrorStack = append(err.ErrorStack, obj)
			i++
		} else {
			break
		}
	}
	return err
}

// Error 实现了官方的错误, 跟官方用法一致, 直接返回错误信息, 不包含错误堆栈信息
func (err *errorString) Error() string {
	return err.msg
}

// ErrorWithStack 跟 func (err *errorString) Error() string 类似, 包含自定义错误信息和错误堆栈信息
func (err *errorString) ErrorWithStack() string {
	var msg = err.msg
	for _, item := range err.ErrorStack {
		msg = fmt.Sprintf("%s, [%s:%v, %s]", msg, item.File, item.Line, item.FuncName)
	}

	return msg
}

// Stack 获取错误的具体堆栈信息
func (err *errorString) Stack() ErrorStack {
	return err.ErrorStack
}
