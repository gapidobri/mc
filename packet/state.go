package packet

type State int

const (
	StateHandshaking State = 0
	StateStatus      State = 1
	StateLogin       State = 2
)
