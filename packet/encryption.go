package packet

type EncryptionReq struct {
	ServerID    string
	PublicKey   LenBytes
	VerifyToken LenBytes
}

func (EncryptionReq) PacketId() int {
	return int(Encryption)
}

type EncryptionRes struct {
	SharedSecret LenBytes
	VerifyToken  LenBytes
}

type LenBytes []byte

func (l *LenBytes) Marshal(w *Writer) error {
	err := w.WriteVarInt(len(*l))
	if err != nil {
		return err
	}

	_, err = w.Write(*l)

	return err
}

func (l *LenBytes) Unmarshal(r *Reader) error {
	length, err := r.ReadVarInt()
	if err != nil {
		return err
	}

	bytes := make([]byte, length)
	_, err = r.Read(bytes)

	*l = bytes

	return err
}
