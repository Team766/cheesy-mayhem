// Code generated by "stringer -type=input"; DO NOT EDIT.

package plc

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[fieldEstop-0]
	_ = x[redEstop1-1]
	_ = x[redAstop1-2]
	_ = x[redEstop2-3]
	_ = x[redAstop2-4]
	_ = x[blueEstop1-5]
	_ = x[blueAstop1-6]
	_ = x[blueEstop2-7]
	_ = x[blueAstop2-8]
	_ = x[fieldRedEstop1-9]
	_ = x[fieldRedEstop2-10]
	_ = x[fieldBlueEstop1-11]
	_ = x[fieldBlueEstop2-12]
	_ = x[redConnected1-13]
	_ = x[redConnected2-14]
	_ = x[redConnected3-15]
	_ = x[blueConnected1-16]
	_ = x[blueConnected2-17]
	_ = x[blueConnected3-18]
	_ = x[redEstop3-19]
	_ = x[redAstop3-20]
	_ = x[blueEstop3-21]
	_ = x[blueAstop3-22]
	_ = x[inputCount-23]
}

const _input_name = "fieldEstopredEstop1redAstop1redEstop2redAstop2blueEstop1blueAstop1blueEstop2blueAstop2fieldRedEstop1fieldRedEstop2fieldBlueEstop1fieldBlueEstop2redConnected1redConnected2redConnected3blueConnected1blueConnected2blueConnected3redEstop3redAstop3blueEstop3blueAstop3inputCount"

var _input_index = [...]uint16{0, 10, 19, 28, 37, 46, 56, 66, 76, 86, 100, 114, 129, 144, 157, 170, 183, 197, 211, 225, 234, 243, 253, 263, 273}

func (i input) String() string {
	if i < 0 || i >= input(len(_input_index)-1) {
		return "input(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _input_name[_input_index[i]:_input_index[i+1]]
}
