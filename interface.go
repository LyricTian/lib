package lib

import (
	"fmt"

	"time"
)

// T Interface instance
func T(v interface{}) Interface {
	var s string
	if v != nil {
		s = fmt.Sprintf("%v", v)
	}
	return inter{v: Str(s)}
}

// Interface Provide a interface{} convert operation
type Interface interface {
	// ToString Convert interface{} to string
	ToString() string
	// ToBytes Convert interface{} to []byte
	ToBytes() []byte
	// ToString Convert interface{} to int64
	ToInt64() int64
	// ToString Convert interface{} to int32
	ToInt32() int32
	// ToString Convert interface{} to int
	ToInt() int
	// ToString Convert interface{} to uint64
	ToUint64() uint64
	// ToString Convert interface{} to uint32
	ToUint32() uint32
	// ToString Convert interface{} to float64
	ToFloat64() float64
	// ToString Convert interface{} to float32
	ToFloat32() float32
	// ToString Convert interface{} to bool
	ToBool() bool
	// ToString Convert interface{} to time,
	// // If error isn't nil return time.Now()
	ToTime(layout string) time.Time
}

type inter struct {
	v Str
}

func (i inter) ToString() string {
	return i.v.ToString()
}

func (i inter) ToBytes() []byte {
	return i.v.ToBytes()
}

func (i inter) ToInt64() int64 {
	return i.v.ToInt64()
}

func (i inter) ToInt32() int32 {
	return i.v.ToInt32()
}

func (i inter) ToInt() int {
	return i.v.ToInt()
}

func (i inter) ToUint64() uint64 {
	return i.v.ToUint64()
}

func (i inter) ToUint32() uint32 {
	return i.v.ToUint32()
}

func (i inter) ToFloat64() float64 {
	return i.v.ToFloat64()
}

func (i inter) ToFloat32() float32 {
	return i.v.ToFloat32()
}

func (i inter) ToBool() bool {
	return i.v.ToBool()
}

func (i inter) ToTime(layout string) time.Time {
	return i.v.ToTime(layout)
}
