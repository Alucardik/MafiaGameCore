package main

import (
	"flag"
	"mafia-core/client"
	"mafia-core/server"
)

var (
	mode = flag.String("mode", "server", "Server or Client mode")
)

func main() {
	flag.Parse()
	if *mode == "server" {
		server.Run(8080)
	} else {
		client.Run()
	}

	//ch := make(chan string)
	//quit := make(chan int)
	//for i := 0; i < 5; i++ {
	//	go func(id int) {
	//		fmt.Printf("Started routine %d\n", id)
	//		select {
	//		case msg := <-ch:
	//			fmt.Printf("%d routine: got %s\n", id, msg)
	//		case <-quit:
	//			return
	//		}
	//	}(i)
	//}
	//
	//for i := 0; i < 5; i++ {
	//	ch <- fmt.Sprintf("Hey %d", i)
	//}
	//
	//fmt.Printf("SEND ALL MSGS\n")
	//time.Sleep(2 * time.Second)
	////quit <- 0
}
