package lib

// B Provide a []byte convert operation
type B []byte

// ToString Convert []byte to string
func (b B) ToString() string {
	return string(b)
}
