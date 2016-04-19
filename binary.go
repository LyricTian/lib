package lib

import (
	"encoding/binary"

	"bytes"
)

// B Provide a []byte conversions
type B []byte

// String Convert []byte to string
func (b B) String() string {
	return string(b)
}

// Buffer Convert []byte to buffer
func (b B) Buffer() *bytes.Buffer {
	return bytes.NewBuffer(b)
}

func (b B) read(v interface{}) error {
	buf := bytes.NewReader(b)
	return binary.Read(buf, binary.LittleEndian, v)
}

// Int64 Convert []byte to int64
func (b B) Int64() (int64, error) {
	var v int64
	err := b.read(&v)
	if err != nil {
		return 0, err
	}
	return v, nil
}

// DefaultInt64 Convert []byte to int64
func (b B) DefaultInt64(defaultVal int64) int64 {
	v, err := b.Int64()
	if err != nil {
		return defaultVal
	}
	return v
}

// Uint64 Convert []byte to uint64
func (b B) Uint64() (uint64, error) {
	var v uint64
	err := b.read(&v)
	if err != nil {
		return 0, err
	}
	return v, nil
}

// DefaultUint64 Convert []byte to uint64
func (b B) DefaultUint64(defaultVal uint64) uint64 {
	v, err := b.Uint64()
	if err != nil {
		return defaultVal
	}
	return v
}

// Float64 Convert []byte to float64
func (b B) Float64() (float64, error) {
	var v float64
	err := b.read(&v)
	if err != nil {
		return 0, err
	}
	return v, nil
}

// DefaultFloat64 Convert []byte to float64
func (b B) DefaultFloat64(defaultVal float64) float64 {
	v, err := b.Float64()
	if err != nil {
		return defaultVal
	}
	return v
}
