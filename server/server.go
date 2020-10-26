package main

import (
	"bufio"
	"fmt"
	"net"
)

type message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
	for {
		conn, _ := ln.Accept()

		conns <- conn
	}
}

func handleClient(client net.Conn, clientid int, msgs chan message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
	var message message
	reader := bufio.NewReader(client)
	for {

		msg, _ := reader.ReadString('\n')
		message.sender = clientid
		message.message = msg
		msgs <- message
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030

	//TODO Create a Listener for TCP connections on the port given above.
	ln, _ := net.Listen("tcp", ":8030")

	//Create a channel for connections
	conns := make(chan net.Conn)

	//Create a channel for messages
	msgs := make(chan message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	ID := 0
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			// - start to asynchronously handle messages from this client
			client := conn
			clients[ID] = conn
			go handleClient(client, ID, msgs)
			ID++

		case msg := <-msgs:
			for i, client := range clients {
				if msg.sender != i {
					fmt.Fprintf(client, msg.message)
				}
			}
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender

		}
	}
}
