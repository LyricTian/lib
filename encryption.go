package lib

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// NewEncryption Create a Encryption instance
// data The encrypted data
func NewEncryption(data []byte) *Encryption {
	return &Encryption{v: data}
}

// Encryption Provide some commonly used encryption function
type Encryption struct {
	v []byte
}

// MD5 md5 encryption
func (e *Encryption) MD5() (string, error) {
	h := md5.New()
	_, err := h.Write(e.v)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Sha1 sha1 encryption
func (e *Encryption) Sha1() (string, error) {
	h := sha1.New()
	_, err := h.Write(e.v)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
