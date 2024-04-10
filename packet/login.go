package packet

type UUID [16]byte

func (u *UUID) Marshal(w *Writer) error {
	_, err := w.Write(u[:])
	return err
}

func (u *UUID) Unmarshal(r *Reader) error {
	var bytes [16]byte
	_, err := r.Read(bytes[:])
	if err != nil {
		return err
	}
	*u = bytes
	return nil
}

type LoginReq struct {
	Name       string
	PlayerUUID UUID
}

func ReadLoginReq(r *Reader) (*LoginReq, error) {
	name, err := r.ReadString()
	if err != nil {
		return nil, err
	}

	playerUUID, err := r.ReadUUID()
	if err != nil {
		return nil, err
	}

	loginReq := &LoginReq{
		Name:       name,
		PlayerUUID: playerUUID,
	}

	return loginReq, nil
}

type (
	LoginRes struct {
		UUID       UUID
		Username   string
		Properties Properties
	}
)

func (LoginRes) PacketId() int {
	return int(LoginSuccess)
}

type (
	Properties []Property
	Property   struct {
		Name      string
		Value     string
		Signature *Signature
	}
)

func (p Properties) Marshal(w *Writer) error {
	err := w.WriteVarInt(len(p))
	if err != nil {
		return err
	}

	for _, property := range p {
		err = Write(w, property)
		if err != nil {
			return err
		}
	}

	return nil
}

type Signature string

func (s *Signature) Marshal(w *Writer) error {
	err := w.WriteBool(w != nil)
	if err != nil {
		return err
	}

	if w != nil {
		return w.WriteString(string(*s))
	}

	return nil
}

func (s *Signature) Unmarshal(r *Reader) error {
	b, err := r.ReadBool()
	if err != nil {
		return err
	}

	if !b {
		return nil
	}

	str, err := r.ReadString()
	if err != nil {
		return err
	}

	*s = Signature(str)

	return nil
}
