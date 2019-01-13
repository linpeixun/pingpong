package server

import (
	"bufio"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/linpeixun/pingpong/message"
	"github.com/soheilhy/cmux"
	"net"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"
)

type PingpongServer struct {
	mu      sync.Mutex
	connMap map[net.Conn]struct{}
}

func (s *PingpongServer) Server(address string) error {
	listener, _ := net.Listen("tcp", address)

	m := cmux.New(listener)

	httpLn := m.Match(cmux.HTTP1Fast())
	rpcLn := m.Match(cmux.Any())

	go s.startHttp(httpLn)
	go m.Serve()

	s.accept(rpcLn)

	return nil
}
func (s *PingpongServer) accept(ln net.Listener) {

	if s.connMap == nil {
		s.connMap = make(map[net.Conn]struct{})
	}
	for {
		conn, err := ln.Accept()
		fmt.Println("rpc")
		if err != nil {
			fmt.Println("tcp err:%s", err)

			continue
		}

		if tc, ok := conn.(*net.TCPConn); ok {
			tc.SetKeepAlive(true)
			tc.SetKeepAlivePeriod(3 * time.Minute)
			tc.SetLinger(10)
		}

		s.mu.Lock()
		s.connMap[conn] = struct{}{}
		s.mu.Unlock()

		go s.connHandler(conn)
	}
}
func (s *PingpongServer) startHttp(ln net.Listener) {
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
func (s *PingpongServer) httpHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("http")
	w.Write([]byte("hello"))

}

const ReaderBuffsize = 1024

func (s *PingpongServer) connHandler(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:%s", err)
		}

		s.mu.Lock()
		delete(s.connMap, conn)
		s.mu.Unlock()

	}()
	reader := bufio.NewReaderSize(conn, ReaderBuffsize)

	for {
		req, err := message.Read(reader)

		if err != nil {
			if !strings.Contains(err.Error(), "closed by the remote host") {
				fmt.Println("read message:%s,%s", err, reflect.TypeOf(err))
			}
			return
		}
		res := req.Clone()
		res.Payload = res.Payload
		conn.Write(res.Encode())

		message.FreeMsg(req)
		message.FreeMsg(req)
	}

}
