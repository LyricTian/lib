package lib

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	Convey("Subject: string convertions test", t, func() {
		Convey("String Test", func() {
			v := "foo"
			So(S(v).String(), ShouldEqual, v)
		})
		Convey("Bytes Test", func() {
			v := "v"
			So(S(v).Bytes(), ShouldHaveLength, 1)
		})
		Convey("Buffer Test", func() {
			v := "1"
			So(S(v).Buffer().Len(), ShouldEqual, 1)
		})
		Convey("Int64 Test", func() {
			v := "100"
			iv, err := S(v).Int64()
			So(err, ShouldBeNil)
			So(iv, ShouldEqual, 100)
			So(S(v).DefaultInt64(0), ShouldEqual, 100)
		})
		Convey("Uint64 Test", func() {
			v := "100.1"
			iv, err := S(v).Uint64()
			So(err, ShouldNotBeNil)
			So(iv, ShouldEqual, 0)
			So(S(v).DefaultUint64(1), ShouldEqual, 1)
		})
		Convey("Float64 Test", func() {
			v := "0.1"
			iv, err := S(v).Float64()
			So(err, ShouldBeNil)
			So(iv, ShouldBeGreaterThan, 0)
			So(S(v).DefaultFloat64(1), ShouldBeGreaterThan, 0)
		})
		Convey("Bool Test", func() {
			v := "true"
			iv, err := S(v).Bool()
			So(err, ShouldBeNil)
			So(iv, ShouldEqual, true)
			So(S(v).DefaultBool(), ShouldEqual, true)
		})
		Convey("Time Test", func() {
			v := time.Now().Format("20060102")
			iv, err := S(v).Time("20060102")
			So(err, ShouldBeNil)
			So(iv.Format("20060102"), ShouldEqual, time.Now().Format("20060102"))
			So(S(v).DefaultTime("20060102"), ShouldHappenOnOrBefore, time.Now())
		})
	})
}
