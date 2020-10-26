package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

//TODO In a continuous loop, read a message from the server and display it.
func read(conn *net.Conn) {
	for {
		reader := bufio.NewReader(*conn)
		msg, _ := reader.ReadString('\n')
		fmt.Printf(msg)
	}
}

func write(conn *net.Conn) {
	stdin := bufio.NewReader(os.Stdin)
	text, _ := stdin.ReadString('\n')
	fmt.Fprintf(*conn, text)
	//TODO Continually get input from the user and send messages to the server.
}

func main() {
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()

	conn, _ := net.Dial("tcp", *addrPtr)

	for {

		go read(&conn)
		write(&conn)
	}

}
