package net

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func Serve(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go func() {
			handleConn(conn)
		}()
	}
}

func handleConn(conn net.Conn) {
	for {
		// 读数据
		bs := make([]byte, 8)

		_, err := conn.Read(bs)

		if err == io.EOF || err == net.ErrClosed || err == io.ErrUnexpectedEOF {
			_ = conn.Close()
			return
		}

		if err != nil {
			continue
		}

		res := handleMsg(bs)
		_, err = conn.Write(res)

		if err == io.EOF || err == net.ErrClosed || err == io.ErrUnexpectedEOF {
			_ = conn.Close()
			return
		}
	}
}

func handleConnV1(conn net.Conn) {
	for {
		bs := make([]byte, 8)
		_, err := conn.Read(bs)

		if err != nil {
			_ = conn.Close()
			return
		}

		res := handleMsg(bs)

		_, err = conn.Write(res)

		if err != nil {
			_ = conn.Close()
			return
		}
	}
}

func handleMsg(bs []byte) []byte {
	return []byte("world")
}

type Server struct {
	addr string
}

func (s *Server) StartAndServer() error {
	listener, err := net.Listen("tcp", s.addr)

	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go func() {
			er := s.handleConn(conn)

			if er != nil {
				_ = conn.Close()
				fmt.Printf("connect error: %v", err)
			}
		}()
	}
}

func (s Server) handleConn(conn net.Conn) error {
	for {
		bs := make([]byte, lenBytes)

		_, err := conn.Read(bs)
		if err != nil {
			return err
		}

		reqBs := make([]byte, binary.BigEndian.Uint64(bs))
		_, err = conn.Read(reqBs)

		if err != nil {
			return err
		}
		res := string(reqBs) + ", from response"

		bs = make([]byte, lenBytes, len(res)+lenBytes)

		binary.BigEndian.PutUint64(bs, uint64(len(res)))
		bs = append(bs, res...)
		if err != nil {
			return err
		}
	}
}
