package e

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("only show a custom errors demo")

	t.Log("msg:",err.Msg)
	t.Log("file:",err.Stack.File)
	t.Log("line:",err.Stack.Line)
	t.Log("func name:",err.Stack.FuncName)
	fmt.Printf("%#v",err)
}


