package json

import (
	"encoding/json"
	"io"
)

// Encode Encode json
func Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// EncodeString Encode json to string
func EncodeString(v interface{}) (string, error) {
	buffer, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

// EncodeReader Encode json to reader
func EncodeReader(v interface{}) io.Reader {
	pr, pw := io.Pipe()
	go func() {
		err := json.NewEncoder(pw).Encode(v)
		if err != nil {
			pw.CloseWithError(err)
		} else {
			pw.Close()
		}
	}()
	return pr
}

// EncodeWriter Encode json to writer
func EncodeWriter(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
