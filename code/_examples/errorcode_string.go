// Code generated by "stringer -type=errorCode"; DO NOT EDIT.

package _examples

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Internal-100000]
}

const _errorCode_name = "Internal"

var _errorCode_index = [...]uint8{0, 8}

func (i errorCode) String() string {
	i -= 100000
	if i < 0 || i >= errorCode(len(_errorCode_index)-1) {
		return "errorCode(" + strconv.FormatInt(int64(i+100000), 10) + ")"
	}
	return _errorCode_name[_errorCode_index[i]:_errorCode_index[i+1]]
}