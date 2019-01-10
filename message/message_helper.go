package message

import "io"

func NewEmptyMessage() *Message {
	header := Header([12]byte{})
	header[0] = magic

	return &Message{Header: &header}
}

func Read(r io.Reader) (*Message, error) {
	msg := NewEmptyMessage()

	err := msg.Decode(r)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
