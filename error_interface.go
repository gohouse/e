package e

// Error 错误包的接口定义
type Error interface {
	error
	Stack() ErrorStackChain
	ErrorWithStack() string
}
