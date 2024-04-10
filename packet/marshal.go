package packet

import (
	"encoding/binary"
	"github.com/pkg/errors"
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
		t := v.Type().Field(i)

		if opt, exists := t.Tag.Lookup("optional"); exists {
			if !v.FieldByName(opt).Bool() {
				continue
			}
		}

		if !f.CanAddr() {
			ptr := reflect.New(f.Type())
			ptr.Elem().Set(f)

			if m, ok := ptr.Interface().(Marshaler); ok {
				err = m.Marshal(w)
				if err != nil {
					return
				}
				continue
			}
		} else {
			if m, ok := f.Addr().Interface().(Marshaler); ok {
				err = m.Marshal(w)
				if err != nil {
					return
				}
				continue
			}
		}

		switch f.Kind() {
		case reflect.Int:
			err = w.WriteVarInt(int(f.Int()))
		case reflect.Int8:
			err = w.WriteByte(byte(f.Int()))
		case reflect.Int16:
			err = binary.Write(w, binary.BigEndian, int16(f.Int()))
		case reflect.Int32:
			err = binary.Write(w, binary.BigEndian, int32(f.Int()))
		case reflect.Int64:
			err = binary.Write(w, binary.BigEndian, f.Int())
		case reflect.Uint8:
			err = binary.Write(w, binary.BigEndian, uint8(f.Uint()))
		case reflect.Uint16:
			err = binary.Write(w, binary.BigEndian, uint16(f.Uint()))
		case reflect.Uint32:
			err = binary.Write(w, binary.BigEndian, uint32(f.Uint()))
		case reflect.Uint64:
			err = binary.Write(w, binary.BigEndian, f.Uint())
		case reflect.String:
			err = w.WriteString(f.String())
		case reflect.Bool:
			err = w.WriteBool(f.Bool())
		case reflect.Struct:
			err = Write(w, f.Interface())
		default:
			err = errors.New("couldn't marshal " + f.Type().String())
		}

		if err != nil {
			return
		}
	}

	err = w.Flush()
	return
}
