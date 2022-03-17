package server

type Player interface {
	SetName(string)
	GetConnectionInfo()
	GetName()
}

type MafiaPlayer struct {
	name                string
	connection          string
	notificationChannel chan notification
}

func (p *MafiaPlayer) SetName(newName string) {
	p.name = newName
}

func (p *MafiaPlayer) GetName() string {
	return p.name
}
