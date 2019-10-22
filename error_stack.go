package e

type IErrorStack interface {
	SetFile(arg string)
	GetFile() string
	SetLine(arg int)
	GetLine() int
	SetFuncName(arg string)
	GetFuncName() string
}
type ErrorStack struct {
	file     string
	line     int
	funcName string
}

func NewErrorStack() *ErrorStack {
	return new(ErrorStack)
}

func (o *ErrorStack) SetFile(arg string) {
	o.file = arg
}

func (o *ErrorStack) GetFile() string {
	return o.file
}

func (o *ErrorStack) SetLine(arg int) {
	o.line = arg
}

func (o *ErrorStack) GetLine() int {
	return o.line
}

func (o *ErrorStack) SetFuncName(arg string) {
	o.funcName = arg
}

func (o *ErrorStack) GetFuncName() string {
	return o.funcName
}
