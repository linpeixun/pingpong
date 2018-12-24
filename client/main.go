package main

import (
	"fmt"
	"github.com/linpeixun/pingpong/message"
	"net"
	"strconv"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8888", 1*time.Second)
	if err != nil {
		fmt.Println("client err:%s", err)

		return
	}
	conn.(*net.TCPConn).SetLinger(10)

	for count := 1; count > 0; {

		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		content := strconv.Itoa(count)
		fmt.Println("%d-%s", count, content)
		m := message.NewMessage()
		m.Data = []byte(content)

		_, err = conn.Write(m.Encode())
		if err != nil {
			fmt.Println(err)
		}

		count--
	}

}
