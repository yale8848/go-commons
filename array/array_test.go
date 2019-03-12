// Create by Yale 2019/3/12 15:28
package array

import (
	"fmt"
	"testing"
)

func TestDivStr(t *testing.T) {

	da:=Div([]string{"1","2","3"},4)

	for  {
		if !da.HasNext() {
			return
		}
		fmt.Println(da.NextStr())
	}

}
