package lib

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInterface(t *testing.T) {
	Convey("Subject: interface{} convertions test", t, func() {
		Convey("String Test", func() {
			v := "foo"
			So(T(v).String(), ShouldEqual, v)
		})
		Convey("Int64 Test", func() {
			v := "100"
			iv, err := T(v).Int64()
			So(err, ShouldBeNil)
			So(iv, ShouldEqual, 100)
			So(T(v).DefaultInt64(0), ShouldEqual, 100)
		})
		Convey("Uint64 Test", func() {
			v := "100.1"
			iv, err := T(v).Uint64()
			So(err, ShouldNotBeNil)
			So(iv, ShouldEqual, 0)
			So(T(v).DefaultUint64(1), ShouldEqual, 1)
		})
		Convey("Float64 Test", func() {
			v := "0.1"
			iv, err := T(v).Float64()
			So(err, ShouldBeNil)
			So(iv, ShouldBeGreaterThan, 0)
			So(T(v).DefaultFloat64(1), ShouldBeGreaterThan, 0)
		})
		Convey("Bool Test", func() {
			v := "true"
			iv, err := T(v).Bool()
			So(err, ShouldBeNil)
			So(iv, ShouldEqual, true)
			So(T(v).DefaultBool(), ShouldEqual, true)
		})
	})
}
