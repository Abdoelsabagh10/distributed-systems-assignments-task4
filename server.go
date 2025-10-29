package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Message structure  
type ChatMessage struct {
	Name    string
	Content string
}

// ChatServer structure 
type ChatServer struct {
	history []string
}

// Function to send message
func (c *ChatServer) SendMessage(msg ChatMessage, reply *[]string) error {
	entry := fmt.Sprintf("[%s]: %s", msg.Name, msg.Content)
	c.history = append(c.history, entry)

	fmt.Println(entry)

	*reply = c.history
	return nil
}

func main() {
	server := new(ChatServer)
	rpc.Register(server)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started on port 1234...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
