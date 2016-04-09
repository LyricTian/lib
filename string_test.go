package lib

import (
	"testing"
	"time"
)

func TestString(t *testing.T) {
	si := "300"
	if Str(si).ToInt() != 300 {
		t.Error("ToInt error.")
	}
	sf := "300.1"
	if Str(sf).ToFloat32() != 300.1 {
		t.Error("ToFloat32 error.")
	}
	st := "20160319"
	t1 := Str(st).ToTime("20060102")
	t2, _ := time.Parse("20060102", st)
	if t1.Unix() != t2.Unix() {
		t.Error("ToTime error.")
	}
}
