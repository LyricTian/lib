package lib_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/LyricTian/lib.v1"
)

func TestRandom(t *testing.T) {
	Convey("Subject: Generate random value test", t, func() {
		random := lib.NewRandom(6)
		Convey("RandomSource Test", func() {
			v := random.Source([]byte("123456abcdef"))
			So(v, ShouldNotEqual, "")
			_, _ = Println("\n RandomSource value:" + v)
		})
		Convey("RandomNumber Test", func() {
			v := random.Number()
			So(v, ShouldNotEqual, "")
			_, _ = Println("\n RandomNumber value:" + v)
		})
		Convey("LowerLetter Test", func() {
			v := random.LowerLetter()
			So(v, ShouldNotEqual, "")
			_, _ = Println("\n LowerLetter value:" + v)
		})
		Convey("UpperLetter Test", func() {
			v := random.UpperLetter()
			So(v, ShouldNotEqual, "")
			_, _ = Println("\n UpperLetter value:" + v)
		})
		Convey("NumberAndLetter Test", func() {
			v := random.NumberAndLetter()
			So(v, ShouldNotEqual, "")
			_, _ = Println("\n NumberAndLetter value:" + v)
		})
	})
}
