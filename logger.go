package e

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogFile(fileNames ...string) HandlerFunc {
	return func(ctx *ErrorContext) {
		var fileName = "./error-statck.log"
		if len(fileNames) > 0 {
			fileName = fileNames[0]
		}
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Println(err)
		}
		//write to the file
		_, err = fmt.Fprint(f, time.Now().Format("[2006-01-02 15:04:05] "), ctx.ErrorWithStack(), "--------------------------------------------------\n")
		if err != nil {
			panic(err.Error())
		}
		f.Close()
	}
}
