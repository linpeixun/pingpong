package main

import "github.com/linpeixun/pingpong"

func main() {
	s := pingpong.PingpongServer{}
	s.Server("127.0.0.1:8888")

}
