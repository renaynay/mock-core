package mock

import "net"

type Server struct {
	Endpoint string
	Listener net.Listener
}

func NewServer(endpoint string) *Server {
	return &Server{Endpoint: endpoint}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.Endpoint)
	if err != nil {
		return err
	}
	s.Listener = listener
	return nil
}

func (s *Server) Stop() error {
	return s.Listener.Close()
}
