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

	wrapped_type := js.ValueOf(string(typ))
	wrapped_callback := js.FuncOf(func(this js.Value, args []js.Value) any {

		fmt.Println(typ + " event fired!")

		if len(args) > 0 {

			event := args[0]

			if !event.IsNull() && !event.IsUndefined() {

				wrapped_event := dom.ToEvent(event)

				fmt.Println(event)
				fmt.Println(wrapped_event)

			}

		}

		return nil

	})
	wrapped_capture := js.ValueOf(capture)

	doc.Value.Call("addEventListener", wrapped_type, wrapped_callback, wrapped_capture)

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
