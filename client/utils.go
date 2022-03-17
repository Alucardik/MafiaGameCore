package client

import (
	"strings"
)

type command uint16

const (
	CONNECT command = iota
	DISCONNECT
	SHOW_PLAYER_LIST
	EXIT
	UNKNOWN
)

func (c command) toString() string {
	switch c {
	case CONNECT:
		return "connect"
	case DISCONNECT:
		return "disconnect"
	case SHOW_PLAYER_LIST:
		return "players"
	case EXIT:
		return "exit"
	default:
		return "undefined"
	}
}

func parseCommand(cmd string) command {
	switch strings.ToLower(cmd) {
	case CONNECT.toString():
		return CONNECT
	case DISCONNECT.toString():
		return DISCONNECT
	case SHOW_PLAYER_LIST.toString():
		return SHOW_PLAYER_LIST
	default:
		return UNKNOWN
	}
}
