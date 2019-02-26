// Create by Yale 2019/2/26 11:21
package cfile

import "os"

func Exist(filePath string) bool {
	_,err:=os.Stat(filePath)
	if err==nil {
		return  true
	}
	return false
}
