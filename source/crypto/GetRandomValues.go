package crypto

import "syscall/js"

func GetRandomValues(length int) []uint8 {

	buffer := make([]byte, length)
	value := js.Global().Get("Uint8Array").New(length)

	js.Global().Get("crypto").Call("getRandomValues", value)
	js.CopyBytesToGo(buffer, value)

	return buffer

}
