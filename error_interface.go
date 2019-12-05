package e

// E 错误包的接口定义
type Error interface {
	error
	Stack() ErrorStack
	ErrorWithStack() string
}
