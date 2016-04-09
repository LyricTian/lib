package mongo

import (
	"testing"
)

func TestIncrID(t *testing.T) {
	handler, err := InitHandler("mongodb://192.168.33.70:27017/test")
	if err != nil {
		t.Error(err)
		return
	}
	defer handler.Session().Close()
	id, err := handler.IncrID("foo")
	if err != nil {
		t.Error(err)
		return
	}
	if id <= 0 {
		t.Error("Get id error:", id)
	}
}
