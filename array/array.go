// Create by Yale 2019/3/12 15:03
package array

import "reflect"

type ArrayDiv interface {
	HasNext()bool
	NextStr()[]string
}

type arrayDiv struct {
	hasNext bool
	len int
	value []interface{}
}

func (ad *arrayDiv)HasNext() bool {

	if len(ad.value) ==0{
		ad.hasNext = false
	}else{
		ad.hasNext = true
	}
	return ad.hasNext
}
func (ad *arrayDiv)next()[]interface{}{
	var nv []interface{}
	if len(ad.value) <= ad.len{
		nv=ad.value[0:len(ad.value)]
		ad.value = []interface{}{}
	}else{
		nv=ad.value[0:ad.len]
		ad.value = ad.value[ad.len:]
	}
	return nv
}
func (ad *arrayDiv)NextStr() []string  {
	nv:=ad.next()
	strs:=make([]string,len(nv))
	for i,v:=range nv{
		strs[i] = v.(string)
	}
	return strs
}
func Div(arrValue interface{},len int)ArrayDiv  {

	vlen:=reflect.ValueOf(arrValue).Len()
	values:=make([]interface{},vlen)
	for i:=0;i<vlen;i++{
		values[i] = reflect.ValueOf(arrValue).Index(i).Interface()
	}
	return &arrayDiv{value:values,len:len}
}
