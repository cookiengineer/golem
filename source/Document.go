package golem

import "golem/dom"
import "syscall/js"

var Document document

func init() {

	doc := js.Global().Get("document")
	html := dom.ToElement(doc.Call("querySelector", "html"))
	head := dom.ToElement(doc.Get("head"))
	body := dom.ToElement(doc.Get("body"))

	Document = document{
		listeners: make(map[dom.EventType][]*dom.EventListener),
		Element:   &html,
		Head:      &head,
		Body:      &body,
		Value:     &doc,
	}

}

type document struct {
	listeners map[dom.EventType][]*dom.EventListener `json:"listeners"`
	Element   *dom.Element                           `json:"element"`
	Head      *dom.Element                           `json:"head"`
	Body      *dom.Element                           `json:"body"`
	Value     *js.Value                              `json:"value"`
}

func (doc *document) AddEventListener(typ dom.EventType, listener dom.EventListener) bool {

	var result bool = false

	wrapped_type := js.ValueOf(string(typ))
	wrapped_callback := js.FuncOf(func(this js.Value, args []js.Value) any {

		if len(args) > 0 {

			event := args[0]

			if !event.IsNull() && !event.IsUndefined() {

				wrapped_event := dom.ToEvent(event)
				listener.Callback(wrapped_event)

			}

		}

		return nil

	})
	wrapped_capture := js.ValueOf(true)

	doc.Value.Call("addEventListener", wrapped_type, wrapped_callback, wrapped_capture)
	listener.Function = &wrapped_callback

	_, ok := doc.listeners[typ]

	if ok == true {
		doc.listeners[typ] = append(doc.listeners[typ], &listener)
		result = true
	} else {
		doc.listeners[typ] = make([]*dom.EventListener, 0)
		doc.listeners[typ] = append(doc.listeners[typ], &listener)
		result = true
	}

	return result

}

func (doc *document) RemoveEventListener(filter_type dom.EventType, listener *dom.EventListener) bool {

	var result bool = false

	if listener != nil {

		listeners, ok := doc.listeners[filter_type]

		if ok == true {

			var index int = -1

			for l := 0; l < len(listeners); l++ {

				if listeners[l].Id == listener.Id {
					index = l
					break
				}

			}

			if index != -1 {

				listener := listeners[index]
				wrapped_type := js.ValueOf(string(filter_type))
				wrapped_callback := *listener.Function
				doc.Value.Call("removeEventListener", wrapped_type, wrapped_callback)

				doc.listeners[filter_type] = append(doc.listeners[filter_type][:index], doc.listeners[filter_type][index+1:]...)

				result = true

			}

		}

	} else {

		listeners, ok := doc.listeners[filter_type]

		if ok == true {

			for l := 0; l < len(listeners); l++ {

				listener := listeners[l]
				wrapped_type := js.ValueOf(string(filter_type))
				wrapped_callback := *listener.Function
				doc.Value.Call("removeEventListener", wrapped_type, wrapped_callback)

			}

			delete(doc.listeners, filter_type)

			result = true

		}

	}

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
