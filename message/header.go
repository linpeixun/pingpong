package message

type Header [12]byte

func (h *Header) CheckMagic() bool {
	return h[0] == magic
}

func (h *Header) Version() byte {
	return h[1]
}

func (h *Header) SetVersion(v byte) {
	h[1] = v
}

func (h *Header) MessageType() MessageType {
	return MessageType(h[2] >> 7)
}

func (h *Header) SetMessageType(mt MessageType) {
	h[2] = h[2] | (byte(mt) << 7)
}
