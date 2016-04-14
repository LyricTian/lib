package lib

import (
	"testing"
)

func TestEncryptionMD5(t *testing.T) {
	encrypt := NewEncryption([]byte("foo"))
	v, err := encrypt.MD5()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
}
