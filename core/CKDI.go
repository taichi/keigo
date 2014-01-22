package core

import (
	"bytes"
)

type CKDI struct {
	SwitchTo States
}

func (cmd *CKDI) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("CKDI")

	cmd.SwitchTo.write(&buffer, NewStates(4),
		func(s State) string { return s.short() })

	return buffer.String()
}
