package network

import (
	"fmt"
	"net"
)

const clientDefaultBufSize = 4096

type TCPClient struct {
	conn net.Conn
}

func NewClient(serverAddress string) (*TCPClient, error) {
	conn, err := net.Dial("tcp", serverAddress)

	if err != nil {
		return nil, err
	}

	return &TCPClient{
		conn: conn,
	}, nil
}

func (c *TCPClient) Send(request []byte) ([]byte, error) {
	_, err := c.conn.Write(request)

	if err != nil {
		return nil, fmt.Errorf("unable to send request: %w", err)
	}

	response := make([]byte, clientDefaultBufSize)
	cnt, err := c.conn.Read(response)

	if err != nil {
		return nil, fmt.Errorf("unable to read response: %w", err)
	}

	return response[:cnt], nil
}

// Close closes TCP client connection
func (c *TCPClient) Close() {
	if c.conn != nil {
		_ = c.conn.Close()
	}
}
