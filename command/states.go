package command

import (
	. "github.com/taichi/keigo/core"
	"strings"
)

func toState(input string) State {
	states := map[string]State{
		"on":      On,
		"off":     Off,
		"enable":  On,
		"disable": Off,
		"e":       On,
		"d":       Off,
		"1":       On,
		"0":       Off,
	}
	return states[strings.ToLower(input)]
}

func toStates(size int, input string) *States {
	mapping := map[rune]State{
		'd': Off,
		'e': On,
		'D': Off,
		'E': On,
		'0': Off,
		'1': On,
	}
	states := *NewStates(size)
	for i, s := range input {
		if x := mapping[s]; x != nil {
			states[i] = x
		}
	}
	return &states
}
