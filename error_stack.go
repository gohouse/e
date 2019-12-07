package e

// StackObject ...
type ErrorStack struct {
	File     string
	Line     int
	FuncName string
}

// ErrorStack 定义结构体
type ErrorStackChain []ErrorStack

// NewErrorStack 初始化
func NewErrorStackChain() ErrorStackChain {
	return ErrorStackChain{}
}
