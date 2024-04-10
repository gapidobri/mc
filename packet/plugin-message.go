package packet

type PluginMessage struct {
	Channel string
	Data    []byte
}

func ReadPluginMessage(r *Reader) (*PluginMessage, error) {
	channel, err := r.ReadString()

	if err != nil {
		return nil, err
	}

	pluginMessage := &PluginMessage{
		Channel: channel,
		Data:    []byte{},
	}

	return pluginMessage, nil
}
