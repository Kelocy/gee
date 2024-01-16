package geecache

// An immutable view of bytes
type ByteView struct {
	b []byte
}

// Return view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// Return a copy of the data nas a byte slice
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// Return data as a string
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
