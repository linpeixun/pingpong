package message

import (
	"bytes"
	"fmt"
	"github.com/imroc/biu"
	"github.com/linpeixun/pingpong/compress"
	"testing"
)

func TestMessage_Encode(t *testing.T) {
	m := GetPooledMsg()

	metaInfo := make(map[string]string)
	metaInfo["key"] = "value"
	m.MetaInfo = metaInfo

	m.SetMessageType(Response)
	m.SetVersion(1)
	m.SetSeq(1222)
	m.SetOneWay(true)
	m.SetHeartbeat(true)
	m.SetCompressType(compress.Nozip)
	m.ServiceMethod = "method"
	m.ServiceId = "serviceId"

	m.Payload = []byte("aaaaaa")

	encodeBytes := m.Encode()

	var buf bytes.Buffer
	buf.Write(encodeBytes)
	fmt.Println(biu.BytesToBinaryString(encodeBytes))

	m2, err := Read(&buf)
	if err != nil {
		t.Fatalf("decode error")
	}
	//m2.ServiceMethod = m.ServiceMethod
	//m2.ServiceId = m.ServiceId

	//fmt.Println(biu.BytesToBinaryString(m2.Encode()))
	if m2.MetaInfo["key"] != m.MetaInfo["key"] {
		t.Errorf("meta error")
	}

	if string(m2.Payload) != "aaaaaa" {
		t.Errorf("payload error")
	}

	if m.ServiceId != m2.ServiceId {
		t.Errorf("ServiceId error,%v   %v", m.ServiceId, m2.ServiceId)
	}

	if m.ServiceMethod != m2.ServiceMethod {
		t.Errorf("ServiceMethod error,%v   %v", m.ServiceMethod, m2.ServiceMethod)
	}

	if m.IsRequest() {
		t.Error("message type error")
	}
	FreeMsg(m)
	FreeMsg(m2)
}
