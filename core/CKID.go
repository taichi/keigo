package core

import (
	"bytes"
	"fmt"
)

type CKID struct {
	SwitchTo State
}

func (cmd *CKID) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("CKID")
	s := cmd.SwitchTo
	if s == On || s == Off {
		fmt.Fprintf(&buffer, " %s", s.long())
	}
	return buffer.String()
}
