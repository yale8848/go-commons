// Create by Yale 2019/2/26 11:57
package cgorun

import "github.com/yale8848/go-commons/sys/crecover"

type Fun func()
func Go(fun Fun,recoverErrMsg crecover.CallMessage)  {
	go func() {
		defer crecover.Recover(recoverErrMsg)
		fun()
	}()
}
