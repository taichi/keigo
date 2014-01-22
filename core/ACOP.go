package core

import (
	"bytes"
	"fmt"
	"github.com/taichi/keigo/log"
)

type ACOP struct {
	Unit   unit
	Wait   Seconds
	Time   Seconds
	States States
}

type unit int

const (
	_                 = iota
	LampOrBuzzer unit = iota
	DirectOutput
)

const min_unit, max_unit = LampOrBuzzer, DirectOutput

func (u unit) write(buffer *bytes.Buffer) {
	if min_unit <= u && u <= max_unit {
		fmt.Fprintf(buffer, " -u %d", u)
	} else {
		log.Debugf("suppressed unit value %d", u)
	}
}

func (cmd *ACOP) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("ACOP")
	cmd.Unit.write(&buffer)

	cmd.States.write(&buffer, NewStates(8),
		func(s State) string { return s.raw() })

	cmd.Wait.write(&buffer, "w")
	cmd.Time.write(&buffer, "t")

	return buffer.String()
}
