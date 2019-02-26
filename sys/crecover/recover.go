// Create by Yale 2018/11/20 16:14
package crecover

import (
	"fmt"
	"runtime"
)

type CallMessage func(message string)

func Recover( callMessage CallMessage) func(){
    fun:= func() {

		if err := recover(); err != nil {
			var stacktrace string
			for i := 1; ; i++ {
				_, f, l, got := runtime.Caller(i)
				if !got {
					break
				}
				stacktrace += fmt.Sprintf("%s:%d\n", f, l)
			}
			logMessage := fmt.Sprintf("Trace: %s\n", err)
			logMessage += fmt.Sprintf("\n%s", stacktrace)

			callMessage(logMessage)
		}
	}

    return  fun

}
