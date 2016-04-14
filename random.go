package lib

import (
	"io"
	"math/rand"
	"time"
)

// NewRandom 创建Random的实例
// l 生成随机码的长度
func NewRandom(l int) *Random {
	return &Random{vl: l}
}

// Random 提供随机码的生成函数
type Random struct {
	vl int
}

// Number 生成只包含数字的随机码
func (this *Random) Number() string {
	source := this.number()
	return this.Source(source)
}

// LowerLetter 生成只包含小写字母的随机码
func (this *Random) LowerLetter() string {
	source := this.lowerLetter()
	return this.Source(source)
}

// UpperLetter 生成只包含大写字母的随机码
func (this *Random) UpperLetter() string {
	source := this.upperLetter()
	return this.Source(source)
}

// NumberAndLetter 生成包含数字和字母(不区分大小写)的随机码
func (this *Random) NumberAndLetter() string {
	source := this.number()
	source = append(source, this.lowerLetter()...)
	source = append(source, this.upperLetter()...)
	return this.Source(source)
}

// Source 从指定的数据源生成随机码
func (this *Random) Source(source []byte) string {
	if len(source) == 0 {
		return ""
	}
	r, w := io.Pipe()
	go func() {
		for i := 0; i < this.vl; i++ {
			rd := rand.New(rand.NewSource(time.Now().UnixNano()))
			val := source[rd.Intn(len(source))]
			w.Write([]byte{val})
		}
		w.Close()
	}()
	var result []byte
	for {
		buf := make([]byte, this.vl)
		n, err := r.Read(buf)
		if err != nil {
			break
		}
		result = append(result, buf[:n]...)
	}
	return string(result)
}

func (this *Random) number() []byte {
	v := make([]byte, 10)
	for i, j := 48, 0; i <= 57; i, j = i+1, j+1 {
		v[j] = byte(i)
	}
	return v
}

func (this *Random) lowerLetter() []byte {
	v := make([]byte, 26)
	for i, j := 97, 0; i < 123; i, j = i+1, j+1 {
		v[j] = byte(i)
	}
	return v
}

func (this *Random) upperLetter() []byte {
	v := make([]byte, 26)
	for i, j := 65, 0; i < 91; i, j = i+1, j+1 {
		v[j] = byte(i)
	}
	return v
}
