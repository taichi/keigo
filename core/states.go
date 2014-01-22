package core

import (
	"bytes"
	"github.com/taichi/keigo/log"
)

type state struct {
	raw, short, long, di func() string
}

var Off = &state{off, short_disable, long_disable, di_off}
var On = &state{on, short_enable, long_enable, di_on}
var Blink = &state{blink, unsupported, long_blink, di_pulse}
var QuickBlink = &state{quickblink, unsupported, unsupported, unsupported}
var Keep = &state{keep, keep, unsupported, unsupported}

var Disable, Enable = Off, On
var TurnOff, TurnOn = Off, On

func off() string        { return "0" }
func on() string         { return "1" }
func blink() string      { return "2" }
func quickblink() string { return "3" }
func keep() string       { return "x" }

func short_disable() string { return "D" }
func long_disable() string  { return "Disable" }
func short_enable() string  { return "E" }
func long_enable() string   { return "Enable" }
func long_blink() string    { return "Blink" }

func di_off() string   { return "TurnOff" }
func di_on() string    { return "TurnOn" }
func di_pulse() string { return "Pulse" }

func unsupported() string {
	log.Panic("Unsupported Operation")
	return ""
}

type State *state
type States []State
type StringerFn func(s State) string

func NewStates(length int) *States {
	newone := make(States, length)
	for i, _ := range newone {
		newone[i] = Keep
	}
	return &newone
}

func (s *States) write(buffer *bytes.Buffer, defaults *States, fn StringerFn) {
	defs := *defaults
	if length := len(*s); 0 < length {
		buffer.WriteRune(' ')
		copy(defs, *s)
		for _, value := range defs {
			if value != nil {
				buffer.WriteString(fn(value))
			} else {
				log.Debugf("suppressed State value %s", value)
			}
		}
	}
}
