package e

// E 错误包的接口定义
type E interface {
	error
	Stack() ErrorStack
	ErrorWithStack() string
	ToError() error
	ToErrorWithStack() error
}
