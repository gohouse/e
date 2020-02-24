// because the former will succeed if err wraps an *os.PathError.
package e

import (
	"fmt"
)

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string, depth ...int) *ErrorContext {
	var dp = NewStores().Getlayer()
	if len(depth) > 0 {
		if depth[0]>dp {
			dp = depth[0]
		}
	}
	err := NewErrorContext()
	err.depth = dp
	err.msg = text
	err.resolveCaller()

	// 开始做中间件处理
	err.handleHandlerFunc()

	return err
}

// NewWithError 初始化错误,并且可以附加原始错误
func NewWithError(msg string, goErr error, depth ...int) *ErrorContext {
	if goErr != nil {
		msg = fmt.Sprint(msg, ":", goErr.Error())
	}
	return New(msg, depth...)
}
