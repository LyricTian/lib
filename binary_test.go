package lib_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"gopkg.in/LyricTian/lib.v1"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBinary(t *testing.T) {
	Convey("Subject: binary convertions test", t, func() {
		var buf bytes.Buffer
		err := binary.Write(&buf, binary.LittleEndian, int64(100))
		So(err, ShouldBeNil)
		v := buf.Bytes()
		Convey("String Test", func() {
			sv := []byte("foo")
			So(lib.B(sv).String(), ShouldEqual, "foo")
		})
		Convey("Buffer Test", func() {
			So(lib.B(v).Buffer().Len(), ShouldBeGreaterThan, 0)
		})
		Convey("Int64 Test", func() {
			iv, err := lib.B(v).Int64()
			So(err, ShouldBeNil)
			So(iv, ShouldEqual, 100)
			So(lib.B(v).DefaultInt64(0), ShouldEqual, 100)
		})
		Convey("Uint64 Test", func() {
			iv, err := lib.B(v).Uint64()
			So(err, ShouldBeNil)
			So(iv, ShouldEqual, 100)
			So(lib.B(v).DefaultUint64(0), ShouldEqual, 100)
		})
		Convey("Float64 Test", func() {
			iv, err := lib.B(v).Float64()
			So(err, ShouldBeNil)
			So(iv, ShouldBeGreaterThan, 0)
			So(lib.B([]byte("10")).DefaultFloat64(10), ShouldEqual, 10)
		})
	})
}
