package lib

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// T Convert interface{} to Interface
func T(v interface{}) Interface {
	return inter{
		v: reflect.Indirect(reflect.ValueOf(v)),
	}
}

// Interface Provide a interface{} conversions
type Interface interface {
	// String Convert interface{} to string
	String() string
	// Int64 Convert interface{} to int64
	Int64() (int64, error)
	// DefaultInt64 Convert interface{} to int64
	DefaultInt64(defaultVal int64) int64
	// Uint64 Convert interface{} to uint64
	Uint64() (uint64, error)
	// DefaultUint64 Convert interface{} to uint64
	DefaultUint64(defaultVal uint64) uint64
	// Float64 Convert interface{} to float64
	Float64() (float64, error)
	// DefaultFloat64 Convert interface{} to float64
	DefaultFloat64(defaultVal float64) float64
	// Bool Convert interface{} to bool
	Bool() (bool, error)
	// DefaultBool Convert interface{} to bool
	DefaultBool() bool
}

type inter struct {
	v reflect.Value
}

func (i inter) kind() reflect.Kind {
	k := i.v.Kind()
	switch {
	case k >= reflect.Int && k <= reflect.Int64:
		return reflect.Int64
	case k >= reflect.Uint && k <= reflect.Uint64:
		return reflect.Uint64
	case k >= reflect.Float32 && k <= reflect.Float64:
		return reflect.Float64
	default:
		return k
	}
}

func (i inter) String() string {
	switch i.kind() {
	case reflect.String:
		return i.v.String()
	case reflect.Int64:
		return strconv.FormatInt(i.v.Int(), 10)
	case reflect.Uint64:
		return strconv.FormatUint(i.v.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(i.v.Float(), 'f', 5, 64)
	case reflect.Bool:
		return strconv.FormatBool(i.v.Bool())
	}
	if !i.v.IsValid() {
		return ""
	}
	return fmt.Sprintf("%v", i.v.Interface())
}

func (i inter) Int64() (int64, error) {
	switch i.kind() {
	case reflect.String:
		return S(i.v.String()).Int64()
	case reflect.Int64:
		return i.v.Int(), nil
	case reflect.Uint64:
		return int64(i.v.Uint()), nil
	case reflect.Float64:
		return int64(i.v.Float()), nil
	}
	if !i.v.IsValid() {
		return 0, nil
	}
	return 0, errors.New("Unknown value type")
}

func (i inter) DefaultInt64(defaultVal int64) int64 {
	v, err := i.Int64()
	if err != nil {
		return defaultVal
	}
	return v
}

func (i inter) Uint64() (uint64, error) {
	switch i.kind() {
	case reflect.Uint64:
		return i.v.Uint(), nil
	case reflect.Int64:
		return uint64(i.v.Int()), nil
	case reflect.Float64:
		return uint64(i.v.Float()), nil
	case reflect.String:
		return S(i.v.String()).Uint64()
	}
	if !i.v.IsValid() {
		return 0, nil
	}
	return 0, errors.New("Unknown value type")
}

func (i inter) DefaultUint64(defaultVal uint64) uint64 {
	v, err := i.Uint64()
	if err != nil {
		return defaultVal
	}
	return v
}

func (i inter) Float64() (float64, error) {
	switch i.kind() {
	case reflect.String:
		return S(i.v.String()).Float64()
	case reflect.Uint64:
		return float64(i.v.Uint()), nil
	case reflect.Int64:
		return float64(i.v.Int()), nil
	case reflect.Float64:
		return i.v.Float(), nil
	}
	if !i.v.IsValid() {
		return 0, nil
	}
	return 0, errors.New("Unknown value type")
}

func (i inter) DefaultFloat64(defaultVal float64) float64 {
	v, err := i.Float64()
	if err != nil {
		return defaultVal
	}
	return v
}

func (i inter) Bool() (bool, error) {
	switch i.kind() {
	case reflect.String:
		return S(i.v.String()).Bool()
	case reflect.Bool:
		return i.v.Bool(), nil
	}
	if !i.v.IsValid() {
		return false, nil
	}
	return false, errors.New("Unknown value type")
}

func (i inter) DefaultBool() bool {
	v, err := i.Bool()
	if err != nil {
		return false
	}
	return v
}
