package message

type MessageType byte

const (
	Request MessageType = iota
	Response
)

func (h *Header) IsRequest() bool {
	return (h[2] >> 7) == 0x0
}
func (h *Header) IsHeartbeat() bool {
	return h[2]<<1>>7 == 0x1
}

func (h *Header) SetHeartbeat(b bool) {
	if b {
		h[2] = h[2] | 0x40
	} else {
		h[2] = h[2] &^ 0x40
	}
}

func (h *Header) SetOneWay(b bool) {
	if b {
		h[2] = h[2] | 0x20
	} else {
		h[2] = h[2] &^ 0x20
	}
}

func (h *Header) IsOneWay() bool {
	return h[2]<<2>>7 == 0x1
}
