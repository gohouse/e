package e

type E interface {
	error
	Stack() ErrorStack
	ErrorWithStack() string
	ToError() error
	ToErrorWithStack() error
}
