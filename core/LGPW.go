package core

import (
	"bytes"
	"fmt"
)

type LGPW struct {
	Password string
}

func (cmd *LGPW) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("LGPW")

	if s := cmd.Password; 0 < len(s) {
		fmt.Fprintf(&buffer, " %s", s)
	}
	return buffer.String()
}
