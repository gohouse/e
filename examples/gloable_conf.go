package main

import (
	"fmt"
	"github.com/gohouse/e"
)

func main() {
	e.NewConfig()
	e.NewConfig().SetLayer(2)
	test_g1()
}
func test_g1()  {
	test_g2()
}
func test_g2()  {
	test_g3()
}
func test_g3()  {
	err := e.New("xxx")
	fmt.Printf("%+v", err.Stack())
}
