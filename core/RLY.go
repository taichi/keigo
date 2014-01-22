package core

import (
	"bytes"
	"fmt"
)

type rly struct {
	SwitchTo State
	Wait     Seconds
	Time     Seconds
}

func writeRLY(n int, r *rly) string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "RLY%d ", n)
	if r.SwitchTo != nil {
		buffer.WriteString(r.SwitchTo.long())
	}

	r.Wait.write(&buffer, "w")
	r.Time.write(&buffer, "t")

	return buffer.String()
}

type RLY1 struct {
	SwitchTo State
	Wait     Seconds
	Time     Seconds
}

func (cmd *RLY1) String() string {
	return writeRLY(1, &rly{cmd.SwitchTo, cmd.Wait, cmd.Time})
}

type RLY2 struct {
	SwitchTo State
	Wait     Seconds
	Time     Seconds
}

func (cmd *RLY2) String() string {
	return writeRLY(2, &rly{cmd.SwitchTo, cmd.Wait, cmd.Time})
}

type RLY3 struct {
	SwitchTo State
	Wait     Seconds
	Time     Seconds
}

func (cmd *RLY3) String() string {
	return writeRLY(3, &rly{cmd.SwitchTo, cmd.Wait, cmd.Time})
}

type RLY4 struct {
	SwitchTo State
	Wait     Seconds
	Time     Seconds
}

func (cmd *RLY4) String() string {
	return writeRLY(4, &rly{cmd.SwitchTo, cmd.Wait, cmd.Time})
}

type RLY5 struct {
	SwitchTo State
	Wait     Seconds
	Time     Seconds
}

func (cmd *RLY5) String() string {
	return writeRLY(5, &rly{cmd.SwitchTo, cmd.Wait, cmd.Time})
}

type RLY6 struct {
	SwitchTo State
	Wait     Seconds
	Time     Seconds
}

func (cmd *RLY6) String() string {
	return writeRLY(6, &rly{cmd.SwitchTo, cmd.Wait, cmd.Time})
}

type RLY7 struct {
	SwitchTo State
	Wait     Seconds
	Time     Seconds
}

func (cmd *RLY7) String() string {
	return writeRLY(7, &rly{cmd.SwitchTo, cmd.Wait, cmd.Time})
}

type RLY8 struct {
	SwitchTo State
	Wait     Seconds
	Time     Seconds
}

func (cmd *RLY8) String() string {
	return writeRLY(8, &rly{cmd.SwitchTo, cmd.Wait, cmd.Time})
}
