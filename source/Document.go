package golem

import "golem/dom"
import "syscall/js"
import "fmt"

var Document document

func init() {

	doc := js.Global().Get("document")

	Document = document{
		listeners: make(map[uint]dom.EventListener),
		Value:     &doc,
	}

}

type document struct {
	listeners map[uint]dom.EventListener `json:"listeners"`
	Value     *js.Value                  `json:"value"`
}

func (doc *document) AddEventListener(typ dom.EventType, listener dom.EventListener, capture bool) bool {

	var result bool = false

	callback := js.FuncOf(func(this js.Value, args []js.Value) any {

		fmt.Println(typ + " event fired!")

		// if len(args) > 0 {

		// 	unwrapped_event := args[0]

		// 	if !unwrapped_event.IsNull() && !unwrapped_event.IsUndefined() {

		// 		event := dom.ToEvent(unwrapped_event)

		// 		fmt.Println(unwrapped_event)
		// 		fmt.Println(event)

		// 	}

		// }

		return nil

	})

	if 1 == 2 {
		fmt.Println(callback)
	}

	doc.Value.Call("addEventListener", js.ValueOf(string(typ)), callback, js.ValueOf(capture))

	// js.FuncOf(func(this js.Value, args []js.Value) any {
	// 	fmt.Println("WTF")
	// 	return nil
	// }), js.ValueOf(true))
	// js.ValueOf(options))

	return result

}

func (doc *document) RemoveEventListener(typ dom.EventType, listener dom.EventListener, usecapture bool) bool {

	var result bool = false

	// TODO

	return result

}

func (doc *document) QuerySelector(query string) *dom.Element {

	var result *dom.Element = nil

	value := doc.Value.Call("querySelector", query)

	if !value.IsNull() && !value.IsUndefined() {
		element := dom.ToElement(value)
		result = &element
	}

	return result

}

func (doc *document) QuerySelectorAll(query string) []*dom.Element {

	var result []*dom.Element

	values := doc.Value.Call("querySelectorAll", query)

	for v := 0; v < values.Length(); v++ {

		value := values.Index(v)

		if !value.IsNull() && !value.IsUndefined() {

			element := dom.ToElement(value)
			result = append(result, &element)

		}

	}

	return result

}
