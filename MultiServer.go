package pingpong

import (
	"bufio"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/soheilhy/cmux"
	"io"
	"net"
	"net/http"
	"sync"
)

type MultiServer struct {
	mu sync.Mutex
}

func (s *MultiServer) Server(address string) error {
	listener, _ := net.Listen("tcp", address)

	m := cmux.New(listener)

	httpLn := m.Match(cmux.HTTP1Fast())
	ln := m.Match(cmux.Any())

	go s.startHttp(httpLn)
	go m.Serve()

	s.accept(ln)

	return nil
}
func (s *MultiServer) accept(ln net.Listener) {
	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Printf("tcp err:%s", err)

			continue
		}

		go s.connHandler(conn)
	}
}
func (s *MultiServer) startHttp(ln net.Listener) {
	router := httprouter.New()

	router.PUT("/*path", s.httpHandler)
	router.POST("/*path", s.httpHandler)
	router.GET("/*path", s.httpHandler)

	httpServer := &http.Server{Handler: router}

	err := httpServer.Serve(ln)
	if err != nil {
		fmt.Println("http server err:%s", err)
	}
}
func (s *MultiServer) httpHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Printf("http")
	w.Write([]byte("hello"))

}

var ReaderBuffsize = 1024

func (s *MultiServer) connHandler(conn net.Conn) {
	reader := bufio.NewReaderSize(conn, ReaderBuffsize)

	var tmp [1024]byte

	io.ReadFull(reader, tmp[:])
	fmt.Printf("receive:%s", string(tmp[:]))
}
