package message

import (
	"fmt"
	"github.com/imroc/biu"
	"github.com/linpeixun/pingpong/compress"
	"testing"
)

func TestHeader_CompressType(t *testing.T) {
	h := Header{}
	h.SetCompressType(compress.Gzip)
	h.SetOneWay(true)
	h.SetMessageType(Response)
	h.SetVersion(1)

	fmt.Println(biu.ByteToBinaryString(h[2]))
	if !(h.CompressType() == compress.Gzip) {
		t.Error("error zip type")
	}
}
