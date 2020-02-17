package e

import (
	"fmt"
	"runtime"
)

// HandlerFunc ...
type HandlerFunc func(ctx *ErrorContext)

// HandlersChain ...
type HandlersChain []HandlerFunc

// ErrorContext 对象
type ErrorContext struct {
	*Stores
	msg         string
	depth       int
	errorStacks ErrorStackChain
}

// NewErrorContext 获取 ErrorContext 对象
func NewErrorContext() *ErrorContext {
	return &ErrorContext{
		Stores:      NewStores(),
		errorStacks: NewErrorStackChain()}
}

// Use handler register
func (ctx *ErrorContext) Use(middlerware ...HandlerFunc) *ErrorContext {
	ctx.handlers = append(ctx.handlers, middlerware...)
	return ctx
}

func (ctx *ErrorContext) handleHandlerFunc() {
	for _, handlerFunc := range ctx.handlers {
		handlerFunc(ctx)
	}
}
func (ctx *ErrorContext) resolveCaller() {
	// 根据设定获取错误堆栈信息
	var i = 0
	for i < ctx.depth {
		// 这里+2是因为,函数调用从封装的函数开始,如果我直接在 main 方法写这个 caller,就是1了
		if funcName, file, line, ok := runtime.Caller(i + 2); ok {
			var obj = ErrorStack{
				File:     file,
				Line:     line,
				FuncName: runtime.FuncForPC(funcName).Name(),
			}
			ctx.errorStacks = append(ctx.errorStacks, obj)
			i++
		} else { // 如果堆栈层数没有这么多,就直接中断,不需要再继续获取了
			break
		}
	}
}

// Error 实现了官方的错误, 跟官方用法一致, 直接返回错误信息, 不包含错误堆栈信息
func (ctx *ErrorContext) Error() string {
	return ctx.msg
}

// ErrorWithStack 跟 func (ctx *ErrorContext) Error() string 类似, 额外包含自定义错误信息和错误堆栈信息
func (ctx *ErrorContext) ErrorWithStack() string {
	var msg = ctx.msg
	for _, item := range ctx.errorStacks {
		//msg = fmt.Sprintf("%s, [%s:%v, %s]", msg, item.File, item.Line, item.FuncName)
		msg = fmt.Sprintf("%s\n%s\n    %s:%v", msg, item.FuncName, item.File, item.Line)
	}

	return fmt.Sprint(msg, "\n")
}

// Stack 获取错误的具体堆栈信息
func (ctx *ErrorContext) Stack() ErrorStackChain {
	return ctx.errorStacks
}
