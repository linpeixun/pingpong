package message

import (
	"fmt"
	"github.com/imroc/biu"
	"testing"
)

func TestHeader_CompressType(t *testing.T) {
	h := Header{}
	h.SetCompressType(Gzip)
	h.SetOneWay(true)
	h.SetMessageType(Response)
	h.SetVersion(1)

	fmt.Println(biu.ByteToBinaryString(h[2]))
	if !(h.CompressType() == Gzip) {
		t.Error("error zip type")
	}
}
