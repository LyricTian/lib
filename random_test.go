package lib_test

import (
	"testing"

	"gopkg.in/LyricTian/lib.v1"
)

func TestRandomSource(t *testing.T) {
	random := lib.NewRandom(6)
	v := random.Source([]byte("123456abcdef"))
	if v == "" {
		t.Error("Generate error")
		return
	}
	t.Log(v)
}

func TestRandomNumber(t *testing.T) {
	random := lib.NewRandom(6)
	v := random.Number()
	if v == "" {
		t.Error("Generate error")
		return
	}
	t.Log(v)
}

func TestRandomLowerLetter(t *testing.T) {
	random := lib.NewRandom(6)
	v := random.LowerLetter()
	if v == "" {
		t.Error("Generate error")
		return
	}
	t.Log(v)
}

func TestRandomUpperLetter(t *testing.T) {
	random := lib.NewRandom(6)
	v := random.UpperLetter()
	if v == "" {
		t.Error("Generate error")
		return
	}
	t.Log(v)
}

func TestRandomNumberAndLetter(t *testing.T) {
	random := lib.NewRandom(6)
	v := random.NumberAndLetter()
	if v == "" {
		t.Error("Generate error")
		return
	}
	t.Log(v)
}
