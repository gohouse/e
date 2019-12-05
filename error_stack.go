package e

type StackObject struct {
	File     string
	Line     int
	FuncName string
}

// ErrorStack 定义结构体
type ErrorStack []StackObject

// NewErrorStack 初始化
func NewErrorStack() ErrorStack {
	return []StackObject{}
}
