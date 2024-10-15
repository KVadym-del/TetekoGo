package src

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func ConnectToServer(ip string, port string) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}(conn)

	fmt.Println("Connected to server. Type your message (or 'quit' to exit):")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if message == "quit" {
			break
		}
		_, err := fmt.Fprintf(conn, message+"\n")
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
		fmt.Print("Server response: ", response)
	}
}
