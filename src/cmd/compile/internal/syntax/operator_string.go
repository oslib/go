// Code generated by "stringer -type Operator -linecomment"; DO NOT EDIT.

package syntax

import "strconv"

const _Operator_name = ":!NOT<-||OR&&AND==!=<<=>>=+-|^*/%&&^<<>>"

var _Operator_index = [...]uint8{0, 1, 2, 5, 7, 9, 11, 13, 16, 18, 20, 21, 23, 24, 26, 27, 28, 29, 30, 31, 32, 33, 34, 36, 38, 40}

func (i Operator) String() string {
	i -= 1
	if i >= Operator(len(_Operator_index)-1) {
		return "Operator(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Operator_name[_Operator_index[i]:_Operator_index[i+1]]
}
