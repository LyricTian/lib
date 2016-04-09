package lib

import (
	"testing"
	"time"
)

func TestInterface(t *testing.T) {
	ti := "300"
	if v := T(ti).ToInt(); v != 300 {
		t.Error("ToInt error:", v)
	}
	tf := "300.1"
	if v := T(tf).ToFloat32(); v != 300.1 {
		t.Error("ToFloat32 error:", v)
	}
	st := "20160319"
	t1 := T(st).ToTime("20060102")
	t2, _ := time.Parse("20060102", st)
	if t1.Unix() != t2.Unix() {
		t.Error("ToTime error.")
	}
}
