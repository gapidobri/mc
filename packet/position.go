package packet

type Position struct {
	X int32
	Z int32
	Y int16
}

func (p *Position) Marshal(w *Writer) error {
	return w.WriteInt64(((int64(p.X) & 0x3FFFFFF) << 38) | ((int64(p.Z) & 0x3FFFFFF) << 12) | (int64(p.Y) & 0xFFF))
}

func (p *Position) Unmarshal(r *Reader) error {
	val, err := r.ReadInt64()
	if err != nil {
		return err
	}

	p.X = int32(val >> 38)
	p.Y = int16(val << 52 >> 52)
	p.Z = int32(val << 26 >> 38)

	return nil
}
