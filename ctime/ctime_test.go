// Create by Yale 2019/1/10 18:07
package ctime

import (
	"fmt"
	"log"
	"testing"
)

func TestGetMonthLastDay(t *testing.T) {

	fmt.Println(GetMonthLastDay("2018-03"))
}
func TestCheckDate(t *testing.T) {


	log.Println(ParseDate("2018-01-09"))

}
