package types

import "syscall/js"
import "strings"

func ToString(value js.Value) string {

	var result string = ""

	typ := value.Type()

	if typ == js.TypeUndefined {

		result = "undefined"

	} else if typ == js.TypeNull {

		result = "null"

	} else if typ == js.TypeBoolean {

		tmp := value.String()

		if strings.HasPrefix(tmp, "<boolean: ") && strings.HasSuffix(tmp, ">") {
			result = tmp[10:len(tmp)-1]
		}

	} else if typ == js.TypeNumber {

		tmp := value.String()

		if strings.HasPrefix(tmp, "<number: ") && strings.HasSuffix(tmp, ">") {
			result = tmp[9:len(tmp)-1]
		}

	} else if typ == js.TypeString {

		result = value.String()

	} else if typ == js.TypeSymbol {

		// TODO: Symbol Support

	} else if typ == js.TypeObject {

		// TODO: Should this be a JSON of an Object?

	} else if typ == js.TypeFunction {

		// TODO: Should this be the Function Body like in JS?

	}

	return result

}
