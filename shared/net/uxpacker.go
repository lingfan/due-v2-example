package net

import (
	"bytes"
	"encoding/binary"
	"io"
	"time"

	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/packet"
)

/**
 * @Description: https://github.com/xhaoh94/UxGame 解析器 客户端采用 LittleEndian
 */
type uxPacker struct {
}

func NewPacker() packet.Packer {
	return &uxPacker{}
}

const defaultSizeBytes = 2
const defaultHeaderBytes = 1
const defaultHeartbeatTimeBytes = 4

// OpType 数据包类型
type OpType uint8

const (
	H_B_S        OpType = 0x01 // 服务器to客户端 心跳
	H_B_R        OpType = 0x02 // 客户端to服务器 心跳
	C_S_C        OpType = 0x03 //
	RPC_REQUIRE  OpType = 0x04 // RPC 请求数据包
	RPC_RESPONSE OpType = 0x05 // RPC 返回数据包
)

func (p *uxPacker) ReadMessage(reader io.Reader) ([]byte, error) {
	buf := make([]byte, defaultSizeBytes)

	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return nil, err
	}

	size := binary.LittleEndian.Uint16(buf)

	log.Infof("ReadMessage buf: %d, %d ", len(buf), size)

	if size == 0 {
		return nil, nil
	}

	data := make([]byte, defaultSizeBytes+size)
	copy(data[:defaultSizeBytes], buf)

	_, err = io.ReadFull(reader, data[defaultSizeBytes:])
	if err != nil {
		return nil, err
	}

	log.Infof("ReadMessage data: %d, %+v ", len(data), data)

	return data, nil
}

func (p *uxPacker) PackMessage(message *packet.Message) ([]byte, error) {

	var (
		err         error
		packageSize = 1 + 4 + 4 + len(message.Buffer)
		buf         = &bytes.Buffer{}
	)
	log.Infof("PackMessage len(message.Buffer):%d, packageSize:%d, message: %#v", len(message.Buffer), packageSize, message)

	buf.Grow(defaultSizeBytes + packageSize)

	err = binary.Write(buf, binary.LittleEndian, uint16(packageSize))
	if err != nil {
		return nil, err
	}

	opType := int8(RPC_RESPONSE)

	if err = binary.Write(buf, binary.LittleEndian, int8(opType)); err != nil {
		return nil, err
	}

	if err = binary.Write(buf, binary.LittleEndian, int32(message.Route)); err != nil {
		return nil, err
	}

	if err = binary.Write(buf, binary.LittleEndian, int32(message.Seq)); err != nil {
		return nil, err
	}

	if err = binary.Write(buf, binary.LittleEndian, message.Buffer); err != nil {
		return nil, err
	}
	log.Infof("PackMessage len(buf.Bytes()):%d, buf.Bytes: %+v", len(buf.Bytes()), buf.Bytes())

	return buf.Bytes(), nil
}

func (p *uxPacker) UnpackMessage(data []byte) (*packet.Message, error) {
	dataLen := len(data)
	var (
		size   uint16
		header uint8
		cmdVal int32
		seqVal int32
		err    error
		ln     = dataLen - defaultSizeBytes - defaultHeaderBytes - 4 - 4 // size(2) + header(1) + seqVal(4)+cmdVal(4)
		reader = bytes.NewReader(data)
	)

	log.Infof("UnpackMessage dataLen: %d ln:%d, data: %+v", dataLen, ln, data)

	if ln < 0 {
		return nil, errors.ErrInvalidMessage
	}

	if err = binary.Read(reader, binary.LittleEndian, &size); err != nil {
		return nil, err
	}
	log.Infof("UnpackMessage size(2): %d", size)

	if err = binary.Read(reader, binary.LittleEndian, &header); err != nil {
		return nil, err
	}
	log.Infof("UnpackMessage header(1): %d", header)

	if err = binary.Read(reader, binary.LittleEndian, &cmdVal); err != nil {
		return nil, err
	}
	log.Infof("UnpackMessage cmdVal(4): %d", cmdVal)

	if err = binary.Read(reader, binary.LittleEndian, &seqVal); err != nil {
		return nil, err
	}
	log.Infof("UnpackMessage seqVal(4): %d", seqVal)

	message := &packet.Message{
		Seq:    seqVal,
		Route:  cmdVal,
		Buffer: make([]byte, ln),
	}

	err = binary.Read(reader, binary.LittleEndian, &message.Buffer)
	if err != nil {
		return nil, err
	}

	log.Infof("UnpackMessage dataLen: %d ln:%d, message: %+v", dataLen, ln, message)

	return message, nil
}

func (p *uxPacker) PackHeartbeat() ([]byte, error) {
	heartbeatBit := uint8(0x02)

	var (
		buf  = &bytes.Buffer{}
		size = defaultHeaderBytes + defaultHeartbeatTimeBytes
	)

	buf.Grow(defaultSizeBytes + size)

	err := binary.Write(buf, binary.LittleEndian, uint16(size))
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, uint8(heartbeatBit))
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, uint32(time.Now().Unix()))
	if err != nil {
		return nil, err
	}
	log.Infof("PackHeartbeat len(buf.Bytes()):%d, buf.Bytes: %+v", len(buf.Bytes()), buf.Bytes())
	return buf.Bytes(), nil
}

func (p *uxPacker) CheckHeartbeat(data []byte) (bool, error) {
	heartbeatBit := uint8(H_B_S)

	log.Infof("CheckHeartbeat data: %d", len(data))

	if len(data) < defaultSizeBytes+defaultHeaderBytes {
		return false, errors.ErrInvalidMessage
	}

	var (
		size   uint16
		header uint8
		reader = bytes.NewReader(data)
	)

	err := binary.Read(reader, binary.LittleEndian, &size)
	log.Infof("CheckHeartbeat size: %d", size)

	if err != nil {
		return false, err
	}

	if uint64(len(data))-uint64(defaultSizeBytes) != uint64(size) {
		return false, errors.ErrInvalidMessage
	}

	err = binary.Read(reader, binary.LittleEndian, &header)
	log.Infof("CheckHeartbeat header: %d", header)

	if err != nil {
		return false, err
	}

	return header&heartbeatBit == heartbeatBit, nil
}
