package slice

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	a := []int{100, 200, 300, 100, 200, 400, 500, 400, 300}
	v := RemoveDuplicates(a)
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
	rc := RandomCapture(data, 30).([]int)
	result := make(map[int]int)
	for i, v := range rc {
		result[v] = i
	}
	if len(result) != 30 {
		t.Fatal("Error:", rc)
	}

	t.Log(rc)
}
