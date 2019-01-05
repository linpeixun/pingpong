package message

import (
	"fmt"
	"github.com/imroc/biu"
	"testing"
)

func TestHeader_MessageType(t *testing.T) {
	h := Header{}
	h.SetMessageType(Request)
	mt := h.MessageType()

	if !(biu.ByteToBinaryString(byte(mt)) == "00000000") {
		t.Error("error")
	}

	h.SetMessageType(Response)
	if !(biu.ByteToBinaryString(byte(h.MessageType())) == "10000000") {
		t.Error("error")
	}
}

func TestHeader_IsHeartbeat(t *testing.T) {
	h := Header{}
	h.SetMessageType(Response)
	h.SetHeartbeat(true)

	fmt.Println(biu.ByteToBinaryString(h[2]))
	if !h.IsHeartbeat() {
		t.Error("must be heartbeat")
	}

	if h.IsRequest() {
		t.Error("must be response")
	}

	h.SetHeartbeat(false)
	fmt.Println(biu.ByteToBinaryString(h[2]))
	if h.IsHeartbeat() {
		t.Error("error heartbeat")
	}
}

func TestHeader_IsOneWay(t *testing.T) {
	h := Header{}

	h.SetMessageType(Response)
	fmt.Println(biu.ByteToBinaryString(h[2]))
	h.SetOneWay(true)
	fmt.Println(biu.ByteToBinaryString(h[2]))

	if h.IsRequest() {
		t.Error("must be request")
	}

	if !h.IsOneWay() {
		t.Error("must be oneway")
	}

	h.SetOneWay(false)
	fmt.Println(biu.ByteToBinaryString(h[2]))
	fmt.Println(biu.ByteToBinaryString(h[2] << 2 >> 5))
	if h.IsOneWay() {
		t.Error("error oneway")
	}
}
