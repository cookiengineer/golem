package timers

import "syscall/js"

type TimeoutCallback func()

func SetInterval(callback TimeoutCallback, millisecond uint) uint {

	var result uint = 0

	wrapped_callback := js.FuncOf(func(this js.Value, args []js.Value) any {
		callback()
	})
	wrapped_delay := js.ValueOf(milliseconds)

	tmp := js.Global().Call("setInterval", wrapped_callback, wrapped_delay)

	if !tmp.IsNull() && !tmp.IsUndefined() {
		result = uint(tmp.Int())
	}

	return result

}
