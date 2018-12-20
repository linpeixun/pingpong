package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client err:%s", err)
	}

	defer conn.Close()

	conn.Write([]byte("1223323"))

}
