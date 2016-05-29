package lib

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncryption(t *testing.T) {
	Convey("Subject: encryption test", t, func() {
		Convey("MD5 Test", func() {
			encrypt := NewEncryption([]byte("foo"))
			v, err := encrypt.MD5()
			So(err, ShouldBeNil)
			So(v, ShouldEqual, "acbd18db4cc2f85cedef654fccc4a4d8")
		})

		Convey("Sha1 Test", func() {
			encrypt := NewEncryption([]byte("foo"))
			v, err := encrypt.Sha1()
			So(err, ShouldBeNil)
			So(v, ShouldEqual, "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33")
		})
	})
}
