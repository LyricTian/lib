package lib

import (
	"bytes"
	"strconv"
	"time"
)

// Str Provide a string convert operation
type Str string

// ToString Convert string to string
func (s Str) ToString() string {
	return string(s)
}

// ToByte Convert string to []byte
func (s Str) ToBytes() []byte {
	return []byte(s.ToString())
}

// ToBuffer Convert string to buffer
func (s Str) ToBuffer() *bytes.Buffer {
	return bytes.NewBufferString(s.ToString())
}

// ToInt64 Convert string to int64
func (s Str) ToInt64() int64 {
	i, err := strconv.ParseInt(s.ToString(), 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// ToInt32 Convert string to int32
func (s Str) ToInt32() int32 {
	return int32(s.ToInt64())
}

// ToInt Convert string to int
func (s Str) ToInt() int {
	return int(s.ToInt32())
}

// ToUint64 Convert string to uint64
func (s Str) ToUint64() uint64 {
	val, err := strconv.ParseUint(s.ToString(), 10, 64)
	if err != nil {
		return 0
	}
	return val
}

// ToUint32 Convert string to uint32
func (s Str) ToUint32() uint32 {
	return uint32(s.ToUint64())
}

// ToFloat64 Convert string to float64
func (s Str) ToFloat64() float64 {
	f, err := strconv.ParseFloat(s.ToString(), 64)
	if err != nil {
		return 0
	}
	return f
}

// ToFloat32 Convert string to float32
func (s Str) ToFloat32() float32 {
	return float32(s.ToFloat64())
}

// ToTime Convert string to time,
// If error isn't nil return time.Now()
func (s Str) ToTime(layout string) time.Time {
	t, err := time.Parse(layout, s.ToString())
	if err != nil {
		return time.Now()
	}
	return t
}

// ToBool Convert string to bool
func (s Str) ToBool() bool {
	v, err := strconv.ParseBool(s.ToString())
	if err != nil {
		return false
	}
	return v
}
