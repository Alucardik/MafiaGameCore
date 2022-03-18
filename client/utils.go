package client

import (
	"strings"
)

type command uint16

const (
	HELP command = iota
	EXIT
	CONNECT
	DISCONNECT
	SHOW_PLAYER_LIST
	VOTE
	END_DAY
	EXPOSE
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
	case VOTE:
		return "vote"
	case EXPOSE:
		return "expose"
	case END_DAY:
		return "skip"
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
	case VOTE.toString():
		return VOTE
	case EXPOSE.toString():
		return EXPOSE
	case END_DAY.toString():
		return END_DAY
	case EXIT.toString():
		return EXIT
	default:
		return UNKNOWN
	}
}
