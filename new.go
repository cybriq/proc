package proc

import "gitlab.com/cybriqsystems/proc/types"

func Item(m *metadata) (t types.Type) {
	switch m.Type() {
	case "bool":
		t = NewBool(m)
	case "duration":
		t = NewDuration(m)
	case "float":
		t = NewFloat(m)
	case "int":
		t = NewInt(m)
	case "list":
		t = NewList(m)
	case "string":
		t = NewString(m)
	case "uint":
		t = NewUint(m)
	default:
		panic("invalid type: '" + m.Type() + "'")
	}
	return
}
