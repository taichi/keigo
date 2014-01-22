package core

import (
	"bytes"
	"fmt"
	"github.com/taichi/keigo/log"
)

type SPOP struct {
	SwitchTo State
	No       SoundNo
	Times    Times
}

type SoundNo uint8

func (s SoundNo) write(buffer *bytes.Buffer) {
	if 0 < s && s < 21 {
		fmt.Fprintf(buffer, "%02d", s)
	} else {
		buffer.WriteString("00")
	}
}

type Times uint

func (t Times) write(buffer *bytes.Buffer) {
	if 0 < t && t < 100 {
		fmt.Fprintf(buffer, "1%02d", t)
	} else {
		buffer.WriteString("000") // repeat
	}
}

func (cmd *SPOP) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("SPOP")

	s := cmd.SwitchTo
	if s == On || s == Off {
		buffer.WriteRune(' ')
		buffer.WriteString(s.raw())
		cmd.No.write(&buffer)
		cmd.Times.write(&buffer)
		buffer.WriteString("00")
	} else {
		log.Debugf("suppressed State %s", s)
	}

	return buffer.String()
}
