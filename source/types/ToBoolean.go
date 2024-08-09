package types

import "syscall/js"

func ToBoolean(value js.Value) bool {

	var result bool = false

	if value.Type() == js.TypeBoolean {
		result = value.Bool()
	}

	return result

}
