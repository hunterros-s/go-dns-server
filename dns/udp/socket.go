package udp

import (
	"fmt"
	"net"

	"github.com/hunterros-s/go-dns-server/dns"
)

type UDPSocketImpl struct {
	conn *net.UDPConn
}

func NewUDPSocket() dns.UDPSocket {
	return &UDPSocketImpl{}
}

func (s *UDPSocketImpl) Bind(server dns.Server) error {
	addr := fmt.Sprintf("%s:%d", server.Address, server.Port)
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}

	s.conn = conn
	return nil
}

func (s *UDPSocketImpl) Unbind() error {
	return s.conn.Close()
}

func (s *UDPSocketImpl) Send_to(data []byte, server dns.Server) error {
	if s.conn == nil {
		return fmt.Errorf("socket not bound")
	}

	addr := fmt.Sprintf("%s:%d", server.Address, server.Port)
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	// Write the data slice to the UDP connection with the specified address
	_, err = s.conn.WriteToUDP(data, udpAddr)
	return err
}

func (s *UDPSocketImpl) Recv_from(buffer []byte) (int, dns.Server, error) {
	if s.conn == nil {
		return 0, dns.Server{}, fmt.Errorf("socket not bound")
	}

	n, remoteAddr, err := s.conn.ReadFromUDP(buffer)
	if err != nil {
		return 0, dns.Server{}, err
	}

	server := dns.Server{
		Address: remoteAddr.IP.String(),
		Port:    remoteAddr.Port,
	}

	return n, server, nil
}
