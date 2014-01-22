package core

import (
	"bytes"
)

type RYOF struct {
	Term term
}

func (cmd *RYOF) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("RYOF")
	cmd.Term.write(&buffer)
	return buffer.String()
}
