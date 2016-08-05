package slice

import (
	"math/rand"
	"reflect"
	"time"
)

// RemoveDuplicates remove duplicate from the slices
func RemoveDuplicates(data interface{}) interface{} {
	iVal := reflect.Indirect(reflect.ValueOf(data))
	if iVal.IsNil() || !iVal.IsValid() || iVal.Kind() != reflect.Slice {
		panic("The unknown slice")
	}
	var el int
	iTyp := reflect.TypeOf(data)
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
	return result.Slice(0, el).Interface()
}

// RandomCapture to obtain a set of random data from the biopsy (don't repeat), returns the length of l slice
func RandomCapture(data interface{}, l int) interface{} {
	dVal := reflect.Indirect(reflect.ValueOf(data))
	if dVal.IsNil() || !dVal.IsValid() || dVal.Kind() != reflect.Slice {
		panic("The unknown slice")
	}
	n := dVal.Len()
	if l > n {
		l = n
	}
	rVal := reflect.MakeSlice(dVal.Type(), l, l)
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		for {
			var exist bool
			v := rd.Intn(n)
			dIndex := dVal.Index(v)
			for j := 0; j < rVal.Len(); j++ {
				if reflect.DeepEqual(dIndex.Interface(), rVal.Index(j).Interface()) {
					exist = true
					break
				}
			}
			if !exist {
				rVal.Index(i).Set(dIndex)
				break
			}
		}
	}
	return rVal.Interface()
}
