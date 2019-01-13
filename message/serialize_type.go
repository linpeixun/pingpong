package message

type SerializeType byte

const (
	None SerializeType = iota
	Json
)

// 0xF0:11110000
func (h *Header) SerializeType() SerializeType {
	return SerializeType((h[3] & 0xF0) >> 4)
}

func (h *Header) SetSerializeType(serializeType SerializeType) {
	h[3] = (h[3] &^ 0xF0) | (byte(serializeType) << 4)
}
