package descape

import (
	"github.com/mattn/go-runewidth"
)

/*
StringWidth calculates the displayed width of given string skipping any terminal
control sequences.
*/
func StringWidth(str string) int {
	// State machine for skipping control sequence ESC [ ... FB (ECMA-48).
	const (
		ESC       = 0x1b
		CSIChar   = '['
		LowestFB  = 0x40
		HighestFB = 0x7e
	)
	const (
		DefaultState = iota
		EscapeState
		ControlState
		ControlExitState
	)
	state := DefaultState
	width := 0
	for _, input := range str {
		switch state {
		case DefaultState, ControlExitState:
			if input == ESC {
				state = EscapeState
			} else {
				state = DefaultState
				width += runewidth.RuneWidth(input)
			}
		case EscapeState:
			if input == CSIChar {
				state = ControlState
			} else {
				state = DefaultState
			}
		case ControlState:
			if LowestFB <= input && input <= HighestFB {
				state = ControlExitState
			}
		}
	}
	return width
}
