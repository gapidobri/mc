package packet

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
)

type Reader struct {
	*bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{bufio.NewReader(r)}
}

const segmentBits = 0x7f
const continueBit = 0x80

func (r *Reader) ReadVarInt() (int, error) {

	var (
		value       byte
		position    int
		currentByte byte
		err         error
	)

	for {
		currentByte, err = r.ReadByte()
		if err != nil {
			return 0, err
		}
		value |= (currentByte & segmentBits) << position

		if currentByte&continueBit == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return 0, errors.New("VarInt is too big")
		}
	}

	return int(value), nil
}

func (r *Reader) ReadString() (string, error) {
	length, err := r.ReadVarInt()
	if err != nil {
		return "", err
	}

	bytes := make([]byte, length)
	_, err = r.Read(bytes)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (r *Reader) ReadUUID() (UUID, error) {
	bytes := make([]byte, 16)
	_, err := r.Read(bytes)
	if err != nil {
		return UUID{}, err
	}

	return UUID(bytes), nil
}

func (r *Reader) ReadUint16() (uint16, error) {
	bytes := make([]byte, 2)
	_, err := r.Read(bytes)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint16(bytes), nil
}

func (r *Reader) ReadInt64() (int64, error) {
	bytes := make([]byte, 8)
	_, err := r.Read(bytes)
	if err != nil {
		return 0, err
	}

	return int64(binary.BigEndian.Uint64(bytes)), nil
}
