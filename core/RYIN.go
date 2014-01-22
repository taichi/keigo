package core

import (
	"bytes"
)

type RYIN struct {
	Term term
}

func (cmd *RYIN) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("RYIN")
	cmd.Term.write(&buffer)
	return buffer.String()
}
