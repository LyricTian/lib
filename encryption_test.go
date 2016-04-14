package lib

import (
	"testing"
)

func TestEncryptionMD5(t *testing.T) {
	encrypt := NewEncryption([]byte("foo"))
	v := encrypt.MD5()
	if v == "" {
		t.Error("Encryption error")
		return
	}
	t.Log(v)
}
