package message

import (
	"encoding/binary"
	"errors"
	"io"
)

type Message struct {
	*Header
	ServiceId     string
	ServiceMethod string
	MetaInfo      map[string]string
	Payload       []byte
	MessageData   []byte
}

func NewMessage() *Message {
	header := Header([12]byte{})
	header[0] = magic

	return &Message{Header: &header}
}

func (m *Message) CheckMagic() bool {
	return m.Header[0] == magic
}
func Read(r io.Reader) (*Message, error) {
	msg := NewMessage()

	_, err := io.ReadFull(r, msg.Header[:])
	if err != nil {
		return nil, err
	}

	if !msg.CheckMagic() {
		return nil, errors.New("error Magic")
	}

	lenData := make([]byte, 4)
	_, err = io.ReadFull(r, lenData[:])

	l := binary.BigEndian.Uint32(lenData)

	dataLen := int(l)

	msg.MessageData = make([]byte, dataLen)
	_, err = io.ReadFull(r, msg.MessageData[:])
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (m *Message) Encode() []byte {
	totalL := 1 + 4 + len(m.MessageData)

	retData := make([]byte, totalL)

	copy(retData[:1], m.Header[:])
	binary.BigEndian.PutUint32(retData[1:5], uint32(len(m.MessageData)))
	copy(retData[5:totalL], m.MessageData)
	return retData
}

const (
	magic byte = 0x08
)
