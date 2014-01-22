package core

import (
	"fmt"
	"github.com/taichi/keigo/log"
	"io"
	"sync"
)

type Command fmt.Stringer

// http://www.isa-j.co.jp/dn1510gl/files/dn1510gl-manual-20130426.pdf
type Session interface {
	io.Closer
	Execute(cmd Command) (string, error)
}

type keigoSession struct {
	sync.Mutex
	config *KeigoConfig
	sock   Socket
}

func Connect(config *KeigoConfig) (Session, error) {
	if s, err := NewSocket(config); err == nil {
		return NewSession(config, s), nil
	} else {
		return nil, err
	}
}

func NewSession(config *KeigoConfig, sock Socket) Session {
	return &keigoSession{config: config, sock: sock}
}

func (s *keigoSession) Close() error {
	s.Lock()
	defer s.Unlock()
	return s.sock.Close()
}

func (s *keigoSession) Execute(cmd Command) (string, error) {
	s.Lock()
	defer s.Unlock()
	cmdstring := cmd.String()
	log.Debugf("> %s", cmdstring)
	for i, times := 0, s.config.Retry+1; i < times; i++ {
		log.Debugf("try %d times", i+1)
		if wRetriable, we := s.sock.WriteLine(cmdstring); we == nil {
			if line, rRetriable, re := s.sock.ReadLine(); re == nil {
				return line, nil // succeed
			} else if rRetriable == false || times < i+1 {
				return "", re
			}
		} else if wRetriable == false || times < i+1 {
			return "", we
		}
		if err := s.sock.Reconnect(); err != nil {
			return "", err
		}
	}
	return "", nil
}
