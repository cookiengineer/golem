package console

import "encoding/json"
import "syscall/js"

func Inspect(instance any) {

	buffer, err := json.MarshalIndent(instance, "", "\t")

	if err == nil {

		value := js.ValueOf(string(buffer))
		object := js.Global().Get("JSON").Call("parse", value)

		js.Global().Get("console").Call("log", object)

	}

}
