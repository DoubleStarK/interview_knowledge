package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		if err != nil {
			return
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			continue
		}
		reader = bufio.NewReader(conn)
		reply, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		fmt.Printf("=== Reply from server: %s\n", reply)
	}
}
