package e

// IErrorStack 定义接口
type IErrorStack interface {
	SetFile(arg string)
	GetFile() string
	SetLine(arg int)
	GetLine() int
	SetFuncName(arg string)
	GetFuncName() string
}

// ErrorStack 定义结构体
type ErrorStack struct {
	file     string
	line     int
	funcName string
}

// NewErrorStack 初始化
func NewErrorStack() *ErrorStack {
	return new(ErrorStack)
}

// SetFile 设置该字段值
func (o *ErrorStack) SetFile(arg string) {
	o.file = arg
}

// GetFile 获取该字段值
func (o *ErrorStack) GetFile() string {
	return o.file
}

// SetLine 设置该字段值
func (o *ErrorStack) SetLine(arg int) {
	o.line = arg
}

// GetLine 获取该字段值
func (o *ErrorStack) GetLine() int {
	return o.line
}

// SetFuncName 设置该字段值
func (o *ErrorStack) SetFuncName(arg string) {
	o.funcName = arg
}

// GetFuncName 获取该字段值
func (o *ErrorStack) GetFuncName() string {
	return o.funcName
}
