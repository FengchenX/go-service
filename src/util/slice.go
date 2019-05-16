package util

import (
	"fmt"
	"reflect"
)

func IsIn(obj interface{}, list interface{}) bool {
	for _, a :=range ToSlice(list){
		if reflect.DeepEqual(obj, a)  {
			return true
		}
	}
	return false
}

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	fmt.Println(v.Kind().String())
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}