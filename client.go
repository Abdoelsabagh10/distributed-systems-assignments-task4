package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

type ChatMessage struct {
	Name    string
	Content string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Welcome, %s! You can start chatting.\n", name)
	fmt.Println("Type 'exit' to quit.")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		var reply []string
		msg := ChatMessage{Name: name, Content: text}

		err = client.Call("ChatServer.SendMessage", msg, &reply)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, m := range reply {
			fmt.Println(m)
		}
		fmt.Println("--------------------\n")
	}
}
