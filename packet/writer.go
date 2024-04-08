package packet

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Writer struct {
	*bufio.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{bufio.NewWriter(w)}
}

func (w *Writer) WriteVarInt(value int) error {
	v := uint(value)
	for {
		if v&^segmentBits == 0 {
			return w.WriteByte(byte(v))
		}

		err := w.WriteByte(byte((v & segmentBits) | continueBit))
		if err != nil {
			return err
		}

		v >>= 7
	}
}

func (w *Writer) WriteString(value string) error {
	v := []byte(value)
	err := w.WriteVarInt(len(v))
	if err != nil {
		return err
	}

	_, err = w.Write(v)
	return err
}

func (w *Writer) WriteInt64(value int64) error {
	bytes := binary.BigEndian.AppendUint64(nil, uint64(value))
	_, err := w.Write(bytes)
	return err
}
