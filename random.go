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
func (rd *Random) Number() string {
	source := rd.number()
	return rd.Source(source)
}

// LowerLetter 生成只包含小写字母的随机码
func (rd *Random) LowerLetter() string {
	source := rd.lowerLetter()
	return rd.Source(source)
}

// UpperLetter 生成只包含大写字母的随机码
func (rd *Random) UpperLetter() string {
	source := rd.upperLetter()
	return rd.Source(source)
}

// NumberAndLetter 生成包含数字和字母(不区分大小写)的随机码
func (rd *Random) NumberAndLetter() string {
	source := rd.number()
	source = append(source, rd.lowerLetter()...)
	source = append(source, rd.upperLetter()...)
	return rd.Source(source)
}

// Source 从指定的数据源生成随机码
func (rd *Random) Source(source []byte) string {
	if len(source) == 0 {
		return ""
	}
	r, w := io.Pipe()
	go func() {
		for i := 0; i < rd.vl; i++ {
			defer func() {
				if err := w.Close(); err != nil {
					panic(err)
				}
			}()
			rd := rand.New(rand.NewSource(time.Now().UnixNano()))
			val := source[rd.Intn(len(source))]
			_, err := w.Write([]byte{val})
			if err != nil {
				panic(err)
			}
		}
	}()
	var result []byte
	for {
		buf := make([]byte, rd.vl)
		n, err := r.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		result = append(result, buf[:n]...)
	}
	return string(result)
}

func (rd *Random) number() []byte {
	v := make([]byte, 10)
	for i, j := 48, 0; i <= 57; i, j = i+1, j+1 {
		v[j] = byte(i)
	}
	return v
}

func (rd *Random) lowerLetter() []byte {
	v := make([]byte, 26)
	for i, j := 97, 0; i < 123; i, j = i+1, j+1 {
		v[j] = byte(i)
	}
	return v
}

func (rd *Random) upperLetter() []byte {
	v := make([]byte, 26)
	for i, j := 65, 0; i < 91; i, j = i+1, j+1 {
		v[j] = byte(i)
	}
	return v
}
