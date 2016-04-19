package lib

import (
	"bytes"
	"strconv"
	"time"
)

// S Provide a string conversions
type S string

// String Convert S to string
func (s S) String() string {
	return string(s)
}

// Bytes Convert string to []byte
func (s S) Bytes() []byte {
	return []byte(s.String())
}

// Buffer Convert string to buffer
func (s S) Buffer() *bytes.Buffer {
	return bytes.NewBufferString(s.String())
}

// Int64 Convert string to int64
func (s S) Int64() (int64, error) {
	return strconv.ParseInt(s.String(), 10, 64)
}

// DefaultInt64 Convert string to int64
func (s S) DefaultInt64(defaultVal int64) int64 {
	v, err := s.Int64()
	if err != nil {
		return defaultVal
	}
	return v
}

// Uint64 Convert string to uint64
func (s S) Uint64() (uint64, error) {
	return strconv.ParseUint(s.String(), 10, 64)
}

// DefaultUint64 Convert string to Uint64
func (s S) DefaultUint64(defaultVal uint64) uint64 {
	v, err := s.Uint64()
	if err != nil {
		return defaultVal
	}
	return v
}

// Float64 Convert string to float64
func (s S) Float64() (float64, error) {
	return strconv.ParseFloat(s.String(), 64)
}

// DefaultFloat64 Convert string to float64
func (s S) DefaultFloat64(defaultVal float64) float64 {
	v, err := s.Float64()
	if err != nil {
		return defaultVal
	}
	return v
}

// Bool Convert string to bool
func (s S) Bool() (bool, error) {
	return strconv.ParseBool(s.String())
}

// DefaultBool Convert string to bool
func (s S) DefaultBool() bool {
	v, err := s.Bool()
	if err != nil {
		return false
	}
	return v
}

// Time Convert string to time
func (s S) Time(layout string) (time.Time, error) {
	return time.Parse(layout, s.String())
}

// DefaultTime Convert string to time,
// If conversion errors,return time.Now()
func (s S) DefaultTime(layout string) time.Time {
	v, err := s.Time(layout)
	if err != nil {
		return time.Now()
	}
	return v
}
