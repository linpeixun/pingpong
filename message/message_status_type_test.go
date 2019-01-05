package message

import (
	"fmt"
	"github.com/imroc/biu"
	"testing"
)

func TestHeader_Error(t *testing.T) {
	h := Header{}
	h.SetMessageType(Request)
	h.SetCompressType(Gzip)
	h.SetHeartbeat(true)
	h.SetOneWay(true)
	h.Error(Normal)

	fmt.Println(biu.ByteToBinaryString(h[2]))

	if h.HasError() {
		t.Error("must be normal")
	}

	h.Error(Error)
	fmt.Println(biu.ByteToBinaryString(h[2]))
	if !h.HasError() {
		t.Error("must be error")
	}

}
