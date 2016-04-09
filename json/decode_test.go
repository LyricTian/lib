package json

import (
	"bytes"
	"testing"
)

func TestDecodeReader2(t *testing.T) {
	var result map[string]string
	br := bytes.NewBufferString(`{"foo":"bar"}`)
	err := DecodeReader2(br, &result)
	if err != nil {
		t.Error(err)
		return
	}
	if result["foo"] != "bar" {
		t.Error("DecodeReader2 error:", result)
	}
}
