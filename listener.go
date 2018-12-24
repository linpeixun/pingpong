package pingpong

import "net"

type ServerListener func(s *PingpongServer, address string) (ln net.Listener, err error)

func tcpServerListener(s *PingpongServer, address string) (ln net.Listener, err error) {
	return nil, nil
}
