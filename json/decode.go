package json

import (
	"encoding/json"
	"io"
)

// DecodeReader Decode json reader
func DecodeReader(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// DecodeReader2 Decode json reader(use pipe)
func DecodeReader2(r io.Reader, v interface{}) error {
	pr, pw := io.Pipe()
	go func() {
		_, err := io.Copy(pw, r)
		if err != nil {
			err = pw.CloseWithError(err)
		} else {
			err = pw.Close()
		}
		if err != nil {
			panic(err)
		}
	}()
	return json.NewDecoder(pr).Decode(v)
}

// DecodeString Decode json string
func DecodeString(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

// Decode Decode json
func Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
