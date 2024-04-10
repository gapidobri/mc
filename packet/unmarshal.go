package packet

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"reflect"
)

type Unmarshaler interface {
	Unmarshal(*Reader) error
}

func Read(r *Reader, value any) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Pointer {
		return errors.New("value is not a pointer")
	}
	v = reflect.Indirect(v)

	if m, ok := v.Addr().Interface().(Unmarshaler); ok {
		return m.Unmarshal(r)
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		t := v.Type().Field(i)

		if opt, exists := t.Tag.Lookup("optional"); exists {
			if !f.FieldByName(opt).Bool() {
				continue
			}
		}

		if m, ok := f.Addr().Interface().(Unmarshaler); ok {
			err := m.Unmarshal(r)
			if err != nil {
				return err
			}
			continue
		}

		switch f.Kind() {
		case reflect.Int:
			val, err := r.ReadVarInt()
			if err != nil {
				return err
			}
			f.SetInt(int64(val))

		case reflect.Int8:
			var val int8
			err := binary.Read(r, binary.BigEndian, &val)
			if err != nil {
				return err
			}
			f.SetInt(int64(val))

		case reflect.Int16:
			var val int16
			err := binary.Read(r, binary.BigEndian, &val)
			if err != nil {
				return err
			}
			f.SetInt(int64(val))

		case reflect.Int32:
			var val int32
			err := binary.Read(r, binary.BigEndian, &val)
			if err != nil {
				return err
			}
			f.SetInt(int64(val))

		case reflect.Int64:
			var val int64
			err := binary.Read(r, binary.BigEndian, &val)
			if err != nil {
				return err
			}
			f.SetInt(val)

		case reflect.Uint8:
			var val uint8
			err := binary.Read(r, binary.BigEndian, &val)
			if err != nil {
				return err
			}
			f.SetUint(uint64(val))

		case reflect.Uint16:
			var val uint16
			err := binary.Read(r, binary.BigEndian, &val)
			if err != nil {
				return err
			}
			f.SetUint(uint64(val))

		case reflect.Uint32:
			var val uint32
			err := binary.Read(r, binary.BigEndian, &val)
			if err != nil {
				return err
			}
			f.SetUint(uint64(val))

		case reflect.Uint64:
			var val uint64
			err := binary.Read(r, binary.BigEndian, &val)
			if err != nil {
				return err
			}
			f.SetUint(val)

		case reflect.String:
			val, err := r.ReadString()
			if err != nil {
				return err
			}
			f.SetString(val)

		case reflect.Bool:
			val, err := r.ReadBool()
			if err != nil {
				return err
			}
			f.SetBool(val)

		case reflect.Struct:
			var val any
			err := Read(r, val)
			if err != nil {
				return err
			}
			f.Set(reflect.ValueOf(val))

		default:
			return errors.New("can't unmarshal field " + f.Type().String() + " in " + v.Type().String())

		}
	}

	return nil
}
