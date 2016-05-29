package mongo

import (
	"testing"
)

func TestIncrID(t *testing.T) {
	handler, err := InitHandlerWithDB("mongodb://admin:123456@45.78.35.157:37017", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer handler.Session().Close()
	id, err := handler.IncrID("foo")
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal("Get id error:", id)
	}
}
