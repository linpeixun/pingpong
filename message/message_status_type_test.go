package message

import (
	"fmt"
	"github.com/imroc/biu"
	"github.com/linpeixun/pingpong/compress"
	"testing"
)

func TestHeader_Error(t *testing.T) {
	h := Header{}
	h.SetMessageType(Request)
	h.SetCompressType(compress.Gzip)
	h.SetHeartbeat(true)
	h.SetOneWay(true)
	h.SetError(Normal)

	fmt.Println(biu.ByteToBinaryString(h[2]))

	if h.HasError() {
		t.Error("must be normal")
	}

	h.SetError(Error)
	fmt.Println(biu.ByteToBinaryString(h[2]))
	if !h.HasError() {
		t.Error("must be error")
	}

}
