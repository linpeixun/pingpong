package message

import (
	"encoding/binary"
	"errors"
	"io"
)

type Message struct {
	*Header
	Data []byte
}

func NewMessage() *Message {
	header := Header([1]byte{})
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

	msg.Data = make([]byte, dataLen)
	_, err = io.ReadFull(r, msg.Data[:])
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (m *Message) Encode() []byte {
	totalL := 1 + 4 + len(m.Data)

	retData := make([]byte, totalL)

	copy(retData[:1], m.Header[:])
	binary.BigEndian.PutUint32(retData[1:5], uint32(len(m.Data)))
	copy(retData[5:totalL], m.Data)
	return retData
}

const (
	magic byte = 0x08
)

type Header [1]byte
