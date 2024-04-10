package play

import "mc/packet"

type LoginReq struct {
	EntityID            int32
	IsHardcore          bool
	DimensionNames      DimensionNames
	MaxPlayers          int
	ViewDistance        int
	SimulationDistance  int
	ReducedDebugInfo    bool
	EnableRespawnScreen bool
	DoLimitedCrafting   bool
	DimensionType       string
	DimensionName       string
	HashedSeed          int64
	GameMode            GameMode
	PreviousGameMode    PreviousGameMode
	IsDebug             bool
	IsFlat              bool
	HasDeathLocation    bool
	DeathDimensionName  *string          `optional:"HasDeathLocation"`
	DeathLocation       *packet.Position `optional:"HasDeathLocation"`
	PortalCooldown      int
}

func (LoginReq) PacketId() int {
	return 0x29
}

type DimensionNames []string

func (d *DimensionNames) Marshal(w *packet.Writer) error {
	err := w.WriteVarInt(len(*d))
	if err != nil {
		return err
	}

	for _, s := range *d {
		err = w.WriteString(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DimensionNames) Unmarshal(r *packet.Reader) error {
	length, err := r.ReadVarInt()
	if err != nil {
		return err
	}

	for i := 0; i < length; i++ {
		var s string
		s, err = r.ReadString()
		if err != nil {
			return err
		}
		*d = append(*d, s)
	}

	return nil
}

type GameMode uint8

const (
	GameModeSurvival  GameMode = 0
	GameModeCreative  GameMode = 1
	GameModeAdventure GameMode = 2
	GameModeSpectator GameMode = 3
)

type PreviousGameMode int8

const (
	PreviousGameModeUndefined PreviousGameMode = -1
	PreviousGameModeSurvival  PreviousGameMode = 0
	PreviousGameModeCreative  PreviousGameMode = 1
	PreviousGameModeAdventure PreviousGameMode = 2
	PreviousGameModeSpectator PreviousGameMode = 3
)
