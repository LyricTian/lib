package lib_test

import (
	"testing"

	"gopkg.in/LyricTian/lib.v1"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInterface(t *testing.T) {
	Convey("Subject: interface{} convertions test", t, func() {
		Convey("String Test", func() {
			v := "foo"
			So(lib.T(v).String(), ShouldEqual, v)
		})
		Convey("Int64 Test", func() {
			v := "100"
			iv, err := lib.T(v).Int64()
			So(err, ShouldBeNil)
			So(iv, ShouldEqual, 100)
			So(lib.T(v).DefaultInt64(0), ShouldEqual, 100)
		})
		Convey("Uint64 Test", func() {
			v := "100.1"
			iv, err := lib.T(v).Uint64()
			So(err, ShouldNotBeNil)
			So(iv, ShouldEqual, 0)
			So(lib.T(v).DefaultUint64(1), ShouldEqual, 1)
		})
		Convey("Float64 Test", func() {
			v := "0.1"
			iv, err := lib.T(v).Float64()
			So(err, ShouldBeNil)
			So(iv, ShouldBeGreaterThan, 0)
			So(lib.T(v).DefaultFloat64(1), ShouldBeGreaterThan, 0)
		})
		Convey("Bool Test", func() {
			v := "true"
			iv, err := lib.T(v).Bool()
			So(err, ShouldBeNil)
			So(iv, ShouldEqual, true)
			So(lib.T(v).DefaultBool(), ShouldEqual, true)
		})
	})
}
