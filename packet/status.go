package packet

import (
	"encoding/json"
)

type StatusRes struct {
	Version            Version     `json:"version"`
	Players            Players     `json:"players"`
	Description        Description `json:"description"`
	Favicon            *string     `json:"favicon"`
	EnforcesSecureChat bool        `json:"enforcesSecureChat"`
	PreviewsChat       bool        `json:"previewsChat"`
}

func (StatusRes) PacketId() int {
	return int(Status)
}

func (s StatusRes) Marshal(w *Writer) error {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return w.WriteString(string(jsonData))
}

type Version struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type Players struct {
	Max    int      `json:"max"`
	Online int      `json:"online"`
	Sample []Sample `json:"sample"`
}

type Sample struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Description struct {
	Text string `json:"text"`
}
