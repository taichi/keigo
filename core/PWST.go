package core

import (
	"bytes"
	"fmt"
)

type PWST struct {
	SwitchTo State
}

func (cmd *PWST) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("PWST")
	s := cmd.SwitchTo
	if s == On || s == Off {
		fmt.Fprintf(&buffer, " %s", s.long())
	}
	return buffer.String()
}
