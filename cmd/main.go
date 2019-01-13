package main

import (
	"github.com/linpeixun/pingpong/server"
)

func main() {
	s := server.PingpongServer{}
	s.Server("127.0.0.1:8888")

}
