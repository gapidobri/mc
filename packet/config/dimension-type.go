package config

type DimensionType struct {
	FixedTime                   *int64  `nbt:"fixed_time"`
	HasSkylight                 bool    `nbt:"has_skylight"`
	HasCeiling                  bool    `nbt:"has_ceiling"`
	Ultrawarm                   bool    `nbt:"ultrawarm"`
	Natural                     bool    `nbt:"natural"`
	CoordinateScale             float64 `nbt:"coordinate_scale"`
	BedWorks                    bool    `nbt:"bed_works"`
	RespawnAnchorWorks          bool    `nbt:"respawn_anchor_works"`
	MinY                        int32   `nbt:"min_y"`
	Height                      int32   `nbt:"height"`
	LocalHeight                 int32   `nbt:"local_height"`
	Infinibum                   string  `nbt:"infinibum"`
	Effects                     string  `nbt:"effects"`
	AmbientLight                float32 `nbt:"ambient_light"`
	PiglinSafe                  bool    `nbt:"piglin_safe"`
	HasRaids                    bool    `nbt:"has_raids"`
	MonsterSpawnLightLevel      int32   `nbt:"monster_spawn_light_level"`
	MonsterSpawnBlockLightLimit int32   `nbt:"monster_spawn_block_light_limit"`
}
