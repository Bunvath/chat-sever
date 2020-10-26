package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(*conn)
	for {

		msg, _ := reader.ReadString('\n')
		fmt.Println("::", msg)

	}

}

func write(conn *net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	stdn := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter the message: ")
		msg, _ := stdn.ReadString('\n')
		if msg == "quit\n" {
			break
		} else {
			fmt.Fprintf(*conn, msg)
		}
	}

}

func main() {

	// Get the server address and port from the commandline arguments.
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")
	go read(&conn)
	write(&conn)
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
}
