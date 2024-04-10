package packet

type PluginMessage struct {
	Channel string
	Data    []byte
}

func (p *PluginMessage) Unmarshal(r *Reader) error {
	channel, err := r.ReadString()
	if err != nil {
		return err
	}
	p.Channel = channel

	switch channel {
	case "minecraft:brand":
		brand, err := r.ReadString()
		if err != nil {
			return err
		}
		p.Data = []byte(brand)
	}

	return nil
}
