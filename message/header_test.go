package message

import "testing"

func TestHeader_CheckMagic(t *testing.T) {
	h := Header{}
	h[0] = 0x11

	if h.CheckMagic() {
		t.Error("must be flase")
	}

	h[0] = magic

	if !h.CheckMagic() {
		t.Error("must be true")
	}
}

func TestHeader_Seq(t *testing.T) {
	h := Header{}
	h.SetSeq(123)

	if h.Seq() != 123 {
		t.Error("error seq")
	}
}

func TestHeader_SetMessageType(t *testing.T) {
	h := Header{}
	h.SetMessageType(Request)

	if h.MessageType() != Request {
		t.Error("must be Request")
	}

	h.SetMessageType(Response)
	if h.MessageType() == Request {
		t.Error("can not be Request")
	}
}

func TestHeader_Version(t *testing.T) {
	h := Header{}
	h.SetVersion(11)

	if h.Version() != 11 {
		t.Error("error version")
	}
}
