// Create by Yale 2019/3/12 16:12
package number

import (
	"strconv"
)

func Decimal64(value float64,prec int) float64 {
	v, _ := strconv.ParseFloat(strconv.FormatFloat(value, 'f', 4, 64), 64)
	return v
}
func Decimal32(value float32,prec int) float32 {
	v, _ := strconv.ParseFloat(strconv.FormatFloat(float64(value), 'f', prec, 32), 32)
	return float32(v)
}

