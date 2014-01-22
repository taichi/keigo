package core

import (
	"bytes"
)

type CKIP struct {
	SwitchTo States
}

func (cmd *CKIP) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("CKIP")

	cmd.SwitchTo.write(&buffer, NewStates(20),
		func(s State) string { return s.short() })

	return buffer.String()
}
