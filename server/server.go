package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	fmt.Println(err)
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	for {
		conn, _ := ln.Accept()
		fmt.Printf("New Client Connected\n")
		conns <- conn
	}
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	reader := bufio.NewReader(client)
	for {
		msg, _ := reader.ReadString('\n')
		sender := clientid
		msg = fmt.Sprintf("%d: %s", sender, msg)
		newMessage := Message{sender, msg}
		msgs <- newMessage
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()
	ln, err := net.Listen("tcp", *portPtr)
	if err != nil {
		handleError(err)
	}
	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)
	//Start accepting connections
	go acceptConns(ln, conns)

	id := 0
	for {
		select {
		case conn := <-conns:
			clientID := id + 1
			id++
			clients[clientID] = conn
			go handleClient(clients[clientID], clientID, msgs)

		case msg := <-msgs:
			for client := range clients {
				if client != msg.sender {
					fmt.Printf("Sending message \"%s\" to client %d\n", msg.message, client)
					go fmt.Fprintf(clients[client], msg.message)
				}
			}
		}
	}
}
