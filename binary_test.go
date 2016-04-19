package lib_test

import (
	"testing"

	"gopkg.in/LyricTian/lib.v1"
)

func TestBinary(t *testing.T) {
	v := []byte("123")
	bv := lib.B(v)
	if v := bv.ToString(); v != "123" {
		t.Error("Convert error:", v)
		return
	}
}
