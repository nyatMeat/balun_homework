package network

import (
	"balun_homework_1/foundation/concurrency"
	"balun_homework_1/foundation/logger"
	"context"
	"errors"
	"fmt"
	"net"
	"time"
)

type Server interface {
	Serve(request string) (string, error)
}

type TCPServer struct {
	ctx       context.Context
	server    Server
	logger    *logger.Logger
	listener  net.Listener
	semaphore *concurrency.Semaphore

	address        string
	idleTimeout    time.Duration
	maxMessageSize int
}

func NewTCPServer(
	ctx context.Context,
	server Server,
	log *logger.Logger,
	address string,
	maxConnections int,
	idleTimeout time.Duration,
	maxMessageSize int,
) *TCPServer {
	return &TCPServer{
		ctx:            ctx,
		server:         server,
		logger:         log,
		address:        address,
		idleTimeout:    idleTimeout,
		maxMessageSize: maxMessageSize,
		semaphore:      concurrency.NewSemaphore(maxConnections),
	}
}

func (ts *TCPServer) Run() error {
	l, err := net.Listen("tcp", ts.address)

	if err != nil {
		return fmt.Errorf("listen error: %w", err)
	}

	ts.listener = l

	for {
		conn, err := ts.listener.Accept()

		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				ts.logger.Error(ts.ctx, "[TCPServer::Run] Run server error", err)
			}

			continue
		}

		err = conn.SetReadDeadline(time.Now().Add(ts.idleTimeout))

		if err != nil {
			ts.logger.Error(ts.ctx, "[TCPServer::Run] Deadline setting error", err)
			return fmt.Errorf("deadline setting error: %w", err)
		}

		ts.semaphore.Acquire()

		go func(conn net.Conn) {
			defer ts.semaphore.Release()

			defer func() {
				if r := recover(); r != nil {
					ts.logger.Error(ts.ctx, "[TCPServer::Run] Error after recover", err)
				}
			}()

			ts.processConnection(conn)
		}(conn)
	}
}

func (ts *TCPServer) processConnection(conn net.Conn) {
	defer func() {
		err := conn.Close()

		if err != nil {
			ts.logger.Error(ts.ctx, "[TCPServer::processConnection] Error connection closing", err)
		}
	}()

	buf := make([]byte, ts.maxMessageSize)

	for {
		requestLen, err := conn.Read(buf)

		if err != nil {
			ts.logger.Error(ts.ctx, "[TCPServer::processConnection] Cannot read request", err)
			break
		}

		if requestLen > ts.maxMessageSize {
			ts.logger.Error(ts.ctx, "[TCPServer::processConnection] Request exceeds max size")
			break
		}

		request := string(buf[:requestLen])
		response, err := ts.server.Serve(request)

		if err != nil {
			response = err.Error()
		}

		ts.logger.Info(ts.ctx, "[TCPServer::processConnection] Send response to client")

		_, err = conn.Write([]byte(response))

		if err != nil {
			ts.logger.Error(ts.ctx, "[TCPServer::processConnection] Write response error", err)
		}
	}
}

func (ts *TCPServer) Stop() {
	err := ts.listener.Close()

	if err != nil {
		ts.logger.Error(ts.ctx, "[TCPServer::stop] Error while stopping server", err)
	}
}
