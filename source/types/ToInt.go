package types

import "syscall/js"
import "strconv"
import "strings"

func ToInt(value js.Value) int {

	var result int = 0

	typ := value.Type()

	if typ == js.TypeUndefined {

		result = 0

	} else if typ == js.TypeNull {

		result = 0

	} else if typ == js.TypeBoolean {

		tmp := value.String()

		if tmp == "<boolean: true>" {
			result = 1
		}

	} else if typ == js.TypeNumber {

		tmp := value.String()

		if strings.HasPrefix(tmp, "<number: ") && strings.HasSuffix(tmp, ">") {

			num, err := strconv.ParseInt(tmp[9:len(tmp)-1], 10, 64)

			if err == nil {
				result = int(num)
			}

		}

	} else if typ == js.TypeString {

		tmp := value.String()

		num, err := strconv.ParseInt(tmp, 10, 64)

		if err == nil {
			result = int(num)
		}

	} else if typ == js.TypeSymbol {

		// TODO: Parse symbol

	} else if typ == js.TypeObject {
		result = 1
	} else if typ == js.TypeFunction {
		result = 1
	}

	return result

}
