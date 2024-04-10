package packet

import (
	"encoding/binary"
	"reflect"
)

type Packet interface {
	PacketId() int
}

type Marshaler interface {
	Marshal(*Writer) error
}

func Write(w *Writer, value any) (err error) {
	v := reflect.Indirect(reflect.ValueOf(value))

	if m, ok := v.Interface().(Marshaler); ok {
		err = m.Marshal(w)
		if err != nil {
			return
		}
		err = w.Flush()
		return
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		if m, ok := f.Interface().(Marshaler); ok {
			err = m.Marshal(w)
			if err != nil {
				return
			}
			continue
		}

		if f.CanInt() {
			if f.Kind() == reflect.Int {
				err = w.WriteVarInt(int(f.Int()))
			} else {
				err = binary.Write(w, binary.BigEndian, f.Int())
			}
		} else {
			switch f.Kind() {
			case reflect.String:
				err = w.WriteString(f.String())
			case reflect.Bool:
				err = w.WriteBool(f.Bool())
			case reflect.Struct:
				err = Write(w, f.Interface())
			default:
			}
		}
		if err != nil {
			return
		}
	}

	err = w.Flush()
	return
}
