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
	conn.(*net.TCPConn).SetLinger(3)

	for count := 100; count > 0; {

		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		content := strconv.Itoa(count)
		m := message.NewEmptyMessage()
		m.ServiceMethod = "method"
		m.ServiceId = "id"
		m.Payload = []byte(content)

		_, err = conn.Write(m.Encode())
		if err != nil {
			fmt.Println(err)
		}

		count--

		res, err := message.Read(conn)

		if err == nil {
			if string(res.Payload) != string(m.Payload) {
				fmt.Println("%d-%s", count, content)
			}
		} else {
			fmt.Println("receive error")
		}
	}

}
