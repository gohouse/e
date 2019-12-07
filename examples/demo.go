package main

import (
	"fmt"
	"github.com/gohouse/e"
	"os"
	"sync"
)

func main() {
	e.NewErrorContext().Use(
		func(ctx *e.ErrorContext) {
			writeFile2(fmt.Sprintf("111 %s 111 \n", ctx.Error()))
		},
		func(ctx *e.ErrorContext) {
			writeFile2(fmt.Sprintf("222 %s 222 \n", ctx.Error()))
		},
		func(ctx *e.ErrorContext) {
			writeFile2(fmt.Sprintf("333 %s 333 \n", ctx.Error()))
		},
	)
	var wg = &sync.WaitGroup{}
	//var hs = e.NewErrorContext()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			act()
			wg.Done()
		}()
	}
	wg.Wait()
}

func act() {
	var err e.Error
	//var fileName2 = "./xxx2.log"
	//var fileName3 = "./xxx3.log"
	// 或者 var err e.Error
	err = test1()

	fmt.Println("error msg:", err.Error())
	fmt.Println("error stack:", err.Stack())
	fmt.Println("error stack:", err.ErrorWithStack())

	//time.Sleep(1*time.Millisecond)
}

func test1() e.Error {
	return test2()
}

func test2() e.Error {
	return test3()
}

func test3() e.Error {
	return e.New("xxxxxxx", 3)
}

//var mu2 sync.RWMutex
func writeFile2(text string) {
	var fileName = "./xxx.log"
	//mu2.Lock()
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	//write to the file
	fmt.Fprintf(f, text)
	f.Close()
	//mu2.Unlock()
}
