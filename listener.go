package pingpong

import "net"

type ServerListener func(s *MultiServer, address string) (ln net.Listener, err error)

func tcpServerListener(s *MultiServer, address string) (ln net.Listener, err error) {
	return nil, nil
}
