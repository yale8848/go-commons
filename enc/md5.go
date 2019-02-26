// Create by Yale 2019/2/26 11:01
package enc

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)
func Md5LowerFile(path string) string {

	t,err:=ioutil.ReadFile(path)
	if err!=nil {
		return ""
	}
	h := md5.New()
	h.Write(t)
	cipherStr := h.Sum(nil)
	return strings.ToLower(fmt.Sprintf("%s", hex.EncodeToString(cipherStr)))
}
func Md5UpperFile(path string) string {

	t,err:=ioutil.ReadFile(path)
	if err!=nil {
		return ""
	}
	h := md5.New()
	h.Write(t)
	cipherStr := h.Sum(nil)
	return strings.ToUpper(fmt.Sprintf("%s", hex.EncodeToString(cipherStr)))
}
func Md5Upper(t string) string {
	h := md5.New()
	h.Write([]byte(t))
	cipherStr := h.Sum(nil)
	return strings.ToUpper(fmt.Sprintf("%s", hex.EncodeToString(cipherStr)))
}
func Md5Lower(t string) string {
	h := md5.New()
	h.Write([]byte(t))
	cipherStr := h.Sum(nil)
	return strings.ToLower(fmt.Sprintf("%s", hex.EncodeToString(cipherStr)))
}

