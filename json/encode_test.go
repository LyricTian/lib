package json

import "testing"

func TestEncodeReader(t *testing.T) {
	data := map[string]string{"foo": "bar"}
	r := EncodeReader(data)
	var result map[string]string
	err := DecodeReader2(r, &result)
	if err != nil {
		t.Error(err)
	}
	if result["foo"] != "bar" {
		t.Error("EncodeReader error")
	}
}
