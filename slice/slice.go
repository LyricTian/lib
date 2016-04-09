package slice

import (
	"errors"
	"reflect"
)

// RemoveDuplicates 去除重复
func RemoveDuplicates(v interface{}) (interface{}, error) {
	iVal := reflect.Indirect(reflect.ValueOf(v))
	if iVal.IsNil() || !iVal.IsValid() || !(iVal.Kind() == reflect.Array || iVal.Kind() == reflect.Slice) {
		return nil, errors.New("i must be a slice or array value")
	}
	var el int
	iTyp := reflect.TypeOf(v)
	l := iVal.Len()
	result := reflect.MakeSlice(iTyp, l, l)
	for i := 0; i < l; i++ {
		var exist bool
		val := iVal.Index(i)
		for j := 0; j < i; j++ {
			if reflect.DeepEqual(val.Interface(), iVal.Index(j).Interface()) {
				exist = true
				break
			}
		}
		if !exist {
			result.Index(el).Set(val)
			el++
		}
	}
	return result.Slice(0, el).Interface(), nil
}
