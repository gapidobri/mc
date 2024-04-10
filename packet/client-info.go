package packet

type ClientInfo struct {
	Locale              string
	ViewDistance        int8
	ChatMode            ChatMode
	ChatColors          bool
	DisplayedSkinParts  uint8
	MainHand            MainHand
	EnableTextFiltering bool
	AllowServerListings bool
}

type ChatMode int

const (
	ChatModeEnabled      ChatMode = 0
	ChatModeCommandsOnly ChatMode = 1
	ChatModeHidden       ChatMode = 2
)

type MainHand int

const (
	MainHandLeft  MainHand = 0
	MainHandRight MainHand = 1
)
