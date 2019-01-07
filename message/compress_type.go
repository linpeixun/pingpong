package message

import "github.com/linpeixun/pingpong/compress"

func (h *Header) SetCompressType(c compress.CompressType) {
	h[2] = h[2] | ((0x1c >> 2 & byte(c)) << 2)
}

func (h *Header) CompressType() compress.CompressType {
	return compress.CompressType((h[2] & 0x1c) >> 2)
}
