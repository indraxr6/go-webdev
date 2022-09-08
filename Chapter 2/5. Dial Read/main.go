package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Connection timeout.")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(input)
		fmt.Fprintf(conn, "You : %s\n", input)
	}
	defer conn.Close()

	fmt.Println("here")

}
