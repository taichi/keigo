package core

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

type Socket interface {
	io.Closer
	Reconnect() error
	WriteLine(string) (bool, error)
	ReadLine() (string, bool, error)
}

type keigoSocket struct {
	r      *bufio.Reader
	w      *bufio.Writer
	conn   net.Conn
	config *KeigoConfig
}

func NewSocket(config *KeigoConfig) (Socket, error) {
	sock := &keigoSocket{
		config: config,
	}
	return sock, sock.prepare()
}

func (s *keigoSocket) prepare() error {
	con, err := net.DialTimeout("tcp", s.config.Address, s.config.Timeout)
	if err == nil {
		s.r = bufio.NewReader(con)
		s.w = bufio.NewWriter(con)
		s.conn = con
	}
	return err
}

func (s *keigoSocket) Reconnect() error {
	s.Close() // suppress error
	return s.prepare()
}

func (s *keigoSocket) Close() error {
	return s.conn.Close()
}

func (s *keigoSocket) write(cmd string) error {
	t := string(s.config.Terminal)
	if _, err := fmt.Fprintf(s.w, "%s%s", cmd, t); err == nil {
		return s.w.Flush()
	} else {
		return err
	}
}

func (s *keigoSocket) WriteLine(cmd string) (bool, error) {
	err := s.write(cmd)
	return isRetriable(err), err
}

func isRetriable(err error) bool {
	if neterr, ok := err.(net.Error); (ok && neterr.Timeout()) || err == io.EOF {
		return true
	}
	return false
}

func (s *keigoSocket) read() (string, error) {
	if line, err := s.r.ReadString(s.config.Terminal); err == nil {
		return strings.TrimRight(line, string(s.config.Terminal)), nil
	} else {
		return "", err
	}
}

func (s *keigoSocket) ReadLine() (line string, retriable bool, err error) {
	line, err = s.read()
	retriable = isRetriable(err)
	return
}
