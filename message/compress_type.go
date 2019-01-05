package message

type CompressType byte

const (
	No CompressType = iota
	Gzip
)

func (h *Header) SetCompressType(c CompressType) {
	h[2] = h[2] | ((0x1c >> 2 & byte(c)) << 2)
}

func (h *Header) CompressType() CompressType {
	return CompressType((h[2] & 0x1c) >> 2)
}
