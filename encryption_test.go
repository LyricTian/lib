package lib_test

import (
	"testing"

	"gopkg.in/LyricTian/lib.v1"
)

func TestEncryptionMD5(t *testing.T) {
	encrypt := lib.NewEncryption([]byte("foo"))
	v, err := encrypt.MD5()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
}
