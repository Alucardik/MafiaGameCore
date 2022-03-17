package client

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"mafia-core/proto"
	"os"
	"strings"
	"time"
)

type client struct {
	dialer      proto.MafiaClient
	id          uint64
	conn        *grpc.ClientConn
	isConnected bool
}

var cl = client{isConnected: false}

func (c *client) checkState() bool {
	if !c.isConnected {
		fmt.Println("Your not connected to a game session, join a server first")
	}

	return c.isConnected
}

func (c *client) Connect() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c.conn = conn
	c.dialer = proto.NewMafiaClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	assignedId, err := cl.dialer.Connect(ctx, &proto.ClientInfo{Name: "Dickie"})
	if err != nil {
		log.Fatalf("couldn't connect to server: %v", err)
	}
	c.id = assignedId.Id
	c.isConnected = true
	fmt.Println("ASSIGNED ID", assignedId.Id)
}

func (c *client) Disconnect() {
	if !c.checkState() {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.dialer.Disconnect(ctx, &proto.ClientId{Id: c.id})
	if err != nil {
		log.Fatalf("Error while Disconnecting: %v", err)
	}

	if err := cl.conn.Close(); err != nil {
	}
}

func (c *client) Subscribe() {
	if !c.checkState() {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := c.dialer.SubscribeToNotifications(ctx, &proto.ClientId{Id: c.id})
	if err != nil {
		log.Fatalf("Subscription Failed")
	}

	for {
		notification, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Receiving notification from server failed, %v", err)
		}
		log.Printf("# Client %d: %s", c.id, notification.Info)
	}
}

func (c *client) ShowPlayersList() {
	if !c.checkState() {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.dialer.ShowPlayersList(ctx, &proto.EmptyMsg{})
	if err != nil {
		log.Fatalf("Couldn't get response from server: %v", err)
	}

	fmt.Println("Players in session:")
	for _, name := range resp.Players {
		fmt.Println(name)
	}
}

// TODO: add CLI via flags

func Run() {
	for reader := bufio.NewReader(os.Stdin); ; {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR", err)
		}

		// TODO: add custom handler for the interrupt signal
		switch parseCommand(strings.TrimSpace(cmd)) {
		case CONNECT:
			cl.Connect()
			go cl.Subscribe()
		case DISCONNECT:
			cl.Disconnect()
		case SHOW_PLAYER_LIST:
			cl.ShowPlayersList()
		case UNKNOWN:
			fmt.Println("Unknown command, print 'help' to see available commands")
		}

	}
}
