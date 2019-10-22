package e

// IErrorStack
type IErrorStack interface {
	SetFile(arg string)
	GetFile() string
	SetLine(arg int)
	GetLine() int
	SetFuncName(arg string)
	GetFuncName() string
}

// ErrorStack
type ErrorStack struct {
	file     string
	line     int
	funcName string
}

// NewErrorStack
func NewErrorStack() *ErrorStack {
	return new(ErrorStack)
}

// SetFile arg type:string
func (o *ErrorStack) SetFile(arg string) {
	o.file = arg
}

// GetFile
func (o *ErrorStack) GetFile() string {
	return o.file
}

// SetLine arg type:int
func (o *ErrorStack) SetLine(arg int) {
	o.line = arg
}

// GetLine
func (o *ErrorStack) GetLine() int {
	return o.line
}

// SetFuncName arg type:string
func (o *ErrorStack) SetFuncName(arg string) {
	o.funcName = arg
}

// GetFuncName
func (o *ErrorStack) GetFuncName() string {
	return o.funcName
}
