package slice_test

import (
	"testing"

	"gopkg.in/LyricTian/lib.v2/slice"
)

func TestRemoveDuplicates(t *testing.T) {
	a := []int{100, 200, 300, 100, 200, 400, 500, 400, 300}
	v := slice.RemoveDuplicates(a)
	va := v.([]int)
	if len(va) != 5 {
		t.Fatal("Error:", va)
	}
	t.Log(va)
}

func TestRandomCapture(t *testing.T) {
	data := make([]int, 100)
	for i := 0; i < len(data); i++ {
		data[i] = i + 1
	}
	v := slice.RandomCapture(data, 10)
	result := v.([]int)
	if len(result) != 10 {
		t.Fatal("Error:", result)
	}
	t.Log(result)
}
