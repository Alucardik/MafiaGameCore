package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mafia-core/proto"
	"net"
	"sync"
)

type server struct {
	proto.UnimplementedMafiaServer
	session MafiaSession
	// TODO: replace with a more suitable identifier
	nextClientId uint64
	sessionStart chan int
	sessionEnd   chan int
	mutex        sync.Mutex
}

func (s *server) broadcast(msg notification) {
	for _, player := range s.session.players {
		player.notificationChannel <- msg
	}
}

func (s *server) Connect(_ context.Context, req *proto.ClientInfo) (*proto.ClientId, error) {
	s.mutex.Lock()
	clientId := s.nextClientId
	// making channels buffered, so that broadcast wouldn't be synchronous
	s.session.players[clientId] = &MafiaPlayer{name: req.Name, notificationChannel: make(chan notification, 10)}
	s.broadcast(notification{CLIENT_CONNECTED, s.session.players[clientId].GetName()})
	s.nextClientId++
	s.mutex.Unlock()
	return &proto.ClientId{Id: clientId}, nil
}

func (s *server) Disconnect(_ context.Context, req *proto.ClientId) (*proto.EmptyMsg, error) {
	s.broadcast(notification{CLIENT_DISCONNECTED, s.session.players[req.Id].GetName()})
	// TODO: check if player receives his own notification and nothing breaks
	close(s.session.players[req.Id].notificationChannel)
	s.mutex.Lock()
	delete(s.session.players, req.Id)
	s.mutex.Unlock()
	return &proto.EmptyMsg{}, nil
}

func (s *server) SubscribeToNotifications(req *proto.ClientId, stream proto.Mafia_SubscribeToNotificationsServer) error {
	for event := range s.session.players[req.Id].notificationChannel {
		switch event.eventType {
		case CLIENT_DISCONNECTED:
			if err := stream.Send(&proto.Notification{Info: "Player " + event.clientName + " disconnected"}); err != nil {
				return err
			}
		case CLIENT_CONNECTED:
			if err := stream.Send(&proto.Notification{Info: "Player " + event.clientName + " connected"}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *server) ShowPlayersList(context.Context, *proto.EmptyMsg) (*proto.PlayersList, error) {
	return &proto.PlayersList{Players: s.session.GetConnectedPlayers()}, nil
}

func (s *server) ObserveSession() {
	select {
	case <-s.sessionStart:
		// maybe start session in a separate goroutine
		s.session.Start()
	case <-s.sessionEnd:
		s.session.End()
	}
}

// TODO: add CLI via flags

func Run(port int) {
	//flag.Parse()
	//listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterMafiaServer(s, &server{
		session:      MafiaSession{players: make(map[uint64]*MafiaPlayer)},
		nextClientId: 0,
		sessionStart: make(chan int),
		sessionEnd:   make(chan int),
	})
	log.Printf("SERVER listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
