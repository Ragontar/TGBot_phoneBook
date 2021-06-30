package main

import (
	"fmt"
	"github.com/Ragontar/TGBot_phoneBook/pkg/handlers"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = handlers.TestHandler(conn)
		if err != nil {
			fmt.Println(err)
		}
	}

}
