// Create by Yale 2019/2/26 10:49
package enc

import (
	"encoding/base64"
	"github.com/wumansgy/goEncrypt"
)


//key 8 bit
func DesEnc(text,key string) string {
	cryptText := goEncrypt.DesCBC_Encrypt([]byte(text), []byte(key))
	return base64.StdEncoding.EncodeToString(cryptText)
}
//key 8 bit
func DesDec(text,key string) string {
	nText:=goEncrypt.DesCBC_Decrypt([]byte(text), []byte(key))
	return string(nText)
}
//key 24 bit
func Des3Enc(text,key string) string  {
	cryptText := goEncrypt.TripleDesEncrypt([]byte(text), []byte(key))
	return base64.StdEncoding.EncodeToString(cryptText)
}
//key 24 bit
func Des3Dec(text,key string) string  {
	nText:=goEncrypt.TripleDesDecrypt([]byte(text), []byte(key))
	return string(nText)
}

