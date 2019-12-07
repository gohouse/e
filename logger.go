package e

import (
	"fmt"
	"os"
)

func Log() HandlerFunc {
	return func(ctx *ErrorContext) {
		var fileName = "./error-statck.log"
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		}
		//write to the file
		_, err = fmt.Fprint(f, "error_msg:", ctx.ErrorWithStack(),"--------------------------------------------------\n")
		if err!=nil {
			panic(err.Error())
		}
		f.Close()
	}
}
