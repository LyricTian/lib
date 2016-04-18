package lib

import (
	"crypto/md5"
	"encoding/hex"
)

// NewEncryption 创建Encryption的实例
// v 需要加密的值
func NewEncryption(v []byte) *Encryption {
	return &Encryption{v: v}
}

// Encryption 提供加密操作
type Encryption struct {
	v []byte
}

// MD5 MD5加密
func (e *Encryption) MD5() (string, error) {
	h := md5.New()
	_, err := h.Write(e.v)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
