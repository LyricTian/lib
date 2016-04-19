package lib_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gopkg.in/LyricTian/lib.v1"
)

func TestEncryption(t *testing.T) {
	Convey("Subject: encryption test", t, func() {
		Convey("MD5 Test", func() {
			encrypt := lib.NewEncryption([]byte("foo"))
			v, err := encrypt.MD5()
			So(err, ShouldBeNil)
			So(v, ShouldNotEqual, "")
			_, _ = Println("\n Md5 Value:" + v)
		})
	})
}
