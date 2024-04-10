package play

type LoginReq struct {
	EntityID            int32
	IsHardcore          bool
	DimensionCount      int
	DimensionNames      []string
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
	PreviousGameMode    GameMode
}

type GameMode uint8
