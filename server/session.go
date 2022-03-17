package server

// PlaySession TODO: extend interface and work with it instead of a struct
type PlaySession interface {
	Start()
	End()
	GetConnectedPlayers() []string
}

type MafiaSession struct {
	players map[uint64]*MafiaPlayer
}

func (ms *MafiaSession) GetConnectedPlayers() []string {
	res := make([]string, len(ms.players))
	for _, player := range ms.players {
		res = append(res, player.GetName())
	}

	return res
}

func (ms *MafiaSession) Start() {
}

func (ms *MafiaSession) End() {
}
