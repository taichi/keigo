package core

import (
	"bytes"
)

type RYOT struct {
	Term     term
	SwitchTo *state
	Wait     Seconds
	Time     Seconds
}

func (cmd *RYOT) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("RYOT")
	cmd.Term.write(&buffer)
	buffer.WriteRune(' ')
	buffer.WriteString(cmd.SwitchTo.di())
	cmd.Wait.write(&buffer, "w")
	cmd.Time.write(&buffer, "t")

	return buffer.String()
}
