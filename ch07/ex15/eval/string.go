package eval

import (
	"fmt"
	"strconv"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'g', -1, 64)
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return fmt.Sprintf(" %s %s %s ", b.x.String(), string(b.op), b.y.String())
}

func (c call) String() string {
	var callStr = c.fn + "("
	for i, arg := range c.args {
		if i > 0 {
			callStr += ", "
		}
		callStr += arg.String()
	}
	callStr += ")"
	return callStr
}

func (m min) String() string {
	var callStr = "["
	for i, arg := range m.args {
		if i > 0 {
			callStr += ", "
		}
		callStr += arg.String()
	}
	callStr += "]"
	return callStr
}
