package slice

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	a := []int{100, 200, 300, 100, 200, 400, 500, 400, 300}
	v, err := RemoveDuplicates(a)
	if err != nil {
		t.Error(err)
		return
	}
	va := v.([]int)
	if len(va) != 5 {
		t.Error("Error:", va)
		return
	}
	t.Log(va)
}
