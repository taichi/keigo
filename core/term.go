package core

import (
	"bytes"
	"fmt"
	"github.com/taichi/keigo/log"
)

type term uint8

func (t term) write(buffer *bytes.Buffer) {
	if 0 < t && t < 5 {
		fmt.Fprintf(buffer, " -t %d", t)
	} else {
		log.Panicf("Unsupported Terminal %d", t)
	}
}
