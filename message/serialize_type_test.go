package message

import (
	"github.com/linpeixun/pingpong/compress"
	"testing"
)

func TestHeader_SerializeType(t *testing.T) {
	h := Header{}
	h.SetMessageType(Request)
	h.SetCompressType(compress.Gzip)
	h.SetHeartbeat(true)
	h.SetOneWay(true)
	h.SetError(Normal)
	h.SetSerializeType(None)

	if h.SerializeType() != None {
		t.Error("st error")
	}

	h.SetSerializeType(Json)
	if h.SerializeType() != Json {
		t.Error("st error")
	}

}
