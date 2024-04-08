package packet

import (
	"bytes"
	"encoding/json"
)

type (
	StatusRes struct {
		Version            Version     `json:"version"`
		Players            Players     `json:"players"`
		Description        Description `json:"description"`
		Favicon            *string     `json:"favicon"`
		EnforcesSecureChat bool        `json:"enforcesSecureChat"`
		PreviewsChat       bool        `json:"previewsChat"`
	}

	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	}

	Players struct {
		Max    int      `json:"max"`
		Online int      `json:"online"`
		Sample []Sample `json:"sample"`
	}

	Sample struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	}

	Description struct {
		Text string `json:"text"`
	}
)

func (s StatusRes) Encode() ([]byte, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	data := bytes.NewBuffer(nil)
	writer := NewWriter(data)
	err = writer.WriteVarInt(int(Status))
	if err != nil {
		return nil, err
	}
	err = writer.WriteString(string(jsonData))
	if err != nil {
		return nil, err
	}

	err = writer.Flush()
	if err != nil {
		return nil, err
	}

	return data.Bytes(), err
}
