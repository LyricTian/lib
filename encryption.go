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
func (this *Encryption) MD5() string {
	h := md5.New()
	var _, err = h.Write(this.v)
	if err == nil {
		return hex.EncodeToString(h.Sum(nil))
	}
	return ""
}
