package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")

	for {
		fmt.Println("Enter: ")
		text, _ := stdin.ReadString('\n')
		fmt.Fprintf(conn, text)
		read(&conn)
	}
}
