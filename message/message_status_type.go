package message

type MessageStatusType byte

const (
	Normal MessageStatusType = iota
	Error
)

func (h *Header) Error(m MessageStatusType) {
	h[2] = h[2] | (byte(m) << 6 >> 6)
}

func (h *Header) HasError() bool {
	return MessageStatusType(h[2]<<6>>6) != Normal
}
