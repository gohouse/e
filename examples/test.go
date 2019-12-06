package main

import (
	"errors"
	"fmt"
	"github.com/gohouse/e"
)
func main() {
	var a,b interface{}
	a = testErr()
	b = testE()
	fmt.Println(a,b)
	//fmt.Println(a.(error).Error())
	//
	//rf := reflect.TypeOf(b)
	//fmt.Println(rf.NumMethod())

	if v,ok := a.(error);ok {
		fmt.Println(v.Error())
	}
	if v,ok := b.(e.Error);ok {
		fmt.Println(v.Error())
	}
}

func testErr() error {
	return errors.New("xxx")
}

func testE() e.Error {
	return e.New("bbb")
}
