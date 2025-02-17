package history

var History history

type history struct {
	Value *js.Value `json:"value"`
}

func init() {

	value := js.Global().Get("window").Get("history")

	History = history{
		Value: &value,
	}

}

func (history *history) Back() {

	history.Value.Call("back")

}

func (history *history) Forward() {

	history.Value.Call("forward")

}

func (history *history) Go(delta int) {

	if direction > 0 {

		value := js.Value(delta)

		fmt.Println(value)

	} else if direction < 0 {
	} else if direction == 0 {
	}

	// TODO

}
