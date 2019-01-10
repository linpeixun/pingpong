package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/linpeixun/pingpong/compress"
	"github.com/linpeixun/pingpong/util"
	"io"
)

const (
	magic byte = 0x08
)

var MaxMessageLength = 0

type Message struct {
	*Header
	ServiceId     string
	ServiceMethod string
	MetaInfo      map[string]string
	Payload       []byte
	MessageData   []byte
}

var zeroHeader Header

func (m *Message) Reset() {
	copy(m.Header[1:], zeroHeader[1:])
	m.ServiceMethod = ""
	m.ServiceId = ""
	m.MetaInfo = nil
	m.Payload = nil
	m.MessageData = nil
}
func (m *Message) CheckMagic() bool {
	return m.Header[0] == magic
}

func (m *Message) Decode(r io.Reader) error {
	_, err := io.ReadFull(r, m.Header[:1])

	if err != nil {
		return err
	}
	if !m.CheckMagic() {
		return fmt.Errorf("wrong magic:%v", m.Header[0])
	}

	_, err = io.ReadFull(r, m.Header[1:])
	if err != nil {
		return err
	}

	messageLenData := poolUint32Data.Get().(*[]byte)
	_, err = io.ReadFull(r, *messageLenData)
	if err != nil {
		poolUint32Data.Put(messageLenData)
		return err
	}
	messageLen := int(binary.BigEndian.Uint32(*messageLenData))
	poolUint32Data.Put(messageLenData)

	if MaxMessageLength > 0 && messageLen > MaxMessageLength {
		return fmt.Errorf("消息太长,%v>%v", messageLen, MaxMessageLength)
	}

	// init message data
	if cap(m.MessageData) >= messageLen {
		m.MessageData = m.MessageData[:messageLen]
	} else {
		m.MessageData = make([]byte, messageLen)
	}

	_, err = io.ReadFull(r, m.MessageData)
	if err != nil {
		return err
	}

	index := 0
	index = m.decodeServiceId(index)
	index = m.decodeServiceMethod(index)
	index = m.decodeMetaInfo(index)
	index, _ = m.decodePayload(index)
	return nil
}

func (m *Message) decodePayload(index int) (int, error) {
	l := int(binary.BigEndian.Uint32(m.MessageData[index : index+4]))

	m.Payload = m.MessageData[index+4 : index+4+l]

	compressor := compress.Compressors[m.CompressType()]
	if compressor == nil {
		return 0, fmt.Errorf("error compress:%v", m.CompressType())
	}
	m.Payload, _ = compressor.Unzip(m.Payload)

	return index + 4 + l, nil
}
func (m *Message) decodeMetaInfo(index int) int {
	l := int(binary.BigEndian.Uint32(m.MessageData[index : index+4]))
	endIndex := index + 4 + l
	for i := index + 4; i < endIndex; {
		keyLen := int(binary.BigEndian.Uint32(m.MessageData[i : i+4]))
		key := util.SliceByteToString(m.MessageData[i+4 : i+4+keyLen])

		i = i + 4 + keyLen

		valueLen := int(binary.BigEndian.Uint32(m.MessageData[i : i+4]))
		value := util.SliceByteToString(m.MessageData[i+4 : i+4+valueLen])

		if m.MetaInfo == nil {
			m.MetaInfo = make(map[string]string)
		}
		m.MetaInfo[key] = value

		i = i + 4 + valueLen
	}

	return endIndex
}
func (m *Message) decodeServiceMethod(index int) int {
	l := binary.BigEndian.Uint32(m.MessageData[index : index+4])
	m.ServiceMethod = util.SliceByteToString(m.MessageData[index+4 : index+4+int(l)])

	return index + 4 + int(l)
}
func (m *Message) decodeServiceId(index int) int {
	l := int(binary.BigEndian.Uint32(m.MessageData[index : index+4]))

	m.ServiceId = util.SliceByteToString(m.MessageData[index+4 : index+4+l])
	return index + 4 + l
}
func (m *Message) Encode() []byte {
	serviceIdData := m.encodeServiceId()
	serviceMehtodData := m.encodeServiceMehtod()
	metaInfoData := m.encodeMetaInfo()
	payloadData := m.encodePayload()

	serviceIdDataLen := len(serviceIdData)
	serviceMehtodDataLen := len(serviceMehtodData)
	metaInfoDataLen := len(metaInfoData)
	payloadDataLen := len(payloadData)

	dataLen := (serviceIdDataLen + serviceMehtodDataLen + 4 + metaInfoDataLen + 4 + payloadDataLen)
	messageLen := len(m.Header) + 4 + dataLen

	retData := make([]byte, messageLen)

	//header
	copy(retData, m.Header[:])
	index := len(m.Header)

	//data length
	binary.BigEndian.PutUint32(retData[index:index+4], uint32(dataLen))
	index += 4

	//serviceId
	copy(retData[index:], serviceIdData)
	index += serviceIdDataLen

	//serviceMethod
	copy(retData[index:], serviceMehtodData)
	index += serviceMehtodDataLen

	//MetaInfo
	binary.BigEndian.PutUint32(retData[index:index+4], uint32(metaInfoDataLen))
	index += 4
	copy(retData[index:], metaInfoData)
	index += metaInfoDataLen

	//payload
	binary.BigEndian.PutUint32(retData[index:index+4], uint32(payloadDataLen))
	index += 4
	copy(retData[index:], payloadData)
	index += payloadDataLen

	return retData
}
func (m *Message) encodePayload() []byte {
	if m.CompressType() != compress.Nozip {
		compressor := compress.Compressors[m.CompressType()]
		if compressor == nil {
			m.SetCompressType(compress.Nozip)
		} else {
			payloadData, err := compressor.Zip(m.Payload)

			if err != nil {
				return payloadData
			} else {
				m.SetCompressType(compress.Nozip)
				return m.Payload
			}
		}
	}
	return m.Payload
}
func (m *Message) encodeServiceId() []byte {
	serviceIdLen := len(m.ServiceId)
	var buf bytes.Buffer
	var lenData = make([]byte, 4)
	binary.BigEndian.PutUint32(lenData, uint32(serviceIdLen))
	buf.Write(lenData)
	buf.Write(util.StringToSliceByte(m.ServiceId))

	return buf.Bytes()
}
func (m *Message) encodeServiceMehtod() []byte {
	serviceMethodLen := len(m.ServiceMethod)
	var buf bytes.Buffer
	var lenData = make([]byte, 4)

	binary.BigEndian.PutUint32(lenData, uint32(serviceMethodLen))
	buf.Write(lenData)
	buf.Write(util.StringToSliceByte(m.ServiceMethod))

	return buf.Bytes()
}

// encode Meta Info
func (m *Message) encodeMetaInfo() []byte {
	if m.MetaInfo == nil {
		return nil
	}

	var buf bytes.Buffer
	var lenData = make([]byte, 4)
	for k, v := range m.MetaInfo {
		binary.BigEndian.PutUint32(lenData, uint32(len(k)))
		buf.Write(lenData)
		buf.Write(util.StringToSliceByte(k))

		binary.BigEndian.PutUint32(lenData, uint32(len(v)))
		buf.Write(lenData)
		buf.Write(util.StringToSliceByte(v))
	}

	return buf.Bytes()
}
