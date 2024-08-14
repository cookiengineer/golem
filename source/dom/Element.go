package dom

import "strings"
import "syscall/js"

type Element struct {
	listeners   map[EventType][]*EventListener `json:"listeners"`
	Id          string                         `json:"id"`
	ClassName   string                         `json:"className"`
	Attributes  map[string]string              `json:"attributes"`
	TagName     string                         `json:"tagName"`
	TextContent string                         `json:"textContent"`
	InnerHTML   string                         `json:"innerHTML"`
	Value       *js.Value                      `json:"value"`
}

func ToElement(value js.Value) Element {

	var element Element

	element.listeners = make(map[EventType][]*EventListener)
	element.Id = value.Get("id").String()
	element.ClassName = value.Get("className").String()
	element.TagName = value.Get("tagName").String()
	element.TextContent = value.Get("textContent").String()
	element.InnerHTML = value.Get("innerHTML").String()
	element.Value = &value

	element.Attributes = make(map[string]string)
	element.RefreshAttributes()

	return element

}

func (element *Element) AddEventListener(typ EventType, listener EventListener) bool {

	var result bool = false

	if element.Id != "" {

		wrapped_type := js.ValueOf(string(typ))
		wrapped_callback := js.FuncOf(func(this js.Value, args []js.Value) any {

			if len(args) > 0 {

				event := args[0]

				if !event.IsNull() && !event.IsUndefined() {

					wrapped_event := ToEvent(event)
					listener.Callback(wrapped_event)

				}

			}

			return nil

		})
		wrapped_capture := js.ValueOf(true)

		element.Value.Call("addEventListener", wrapped_type, wrapped_callback, wrapped_capture)

		_, ok := element.listeners[typ]

		if ok == true {
			element.listeners[typ] = append(element.listeners[typ], &listener)
			result = true
		} else {
			element.listeners[typ] = make([]*EventListener, 0)
			element.listeners[typ] = append(element.listeners[typ], &listener)
			result = true
		}

	}

	return result

}

func (element *Element) RemoveEventListener(filter_type EventType, listener *EventListener) bool {

	var result bool = false

	if element.Id != "" {

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

	}

}

func (element *Element) RefreshAttributes() {

	attributes := element.Value.Call("getAttributeNames")

	if !attributes.IsNull() && !attributes.IsUndefined() {

		for key, _ := range element.Attributes {
			delete(element.Attributes, key)
		}

		for a := 0; a < attributes.Length(); a++ {

			name := attributes.Index(a)
			value := element.Value.Call("getAttribute", name)

			if !value.IsNull() {
				element.Attributes[name.String()] = value.String()
			}

		}

	}

}

func (element *Element) GetAttribute(name string) string {

	var value string

	check := validateXMLName(name)

	if check == nil {

		tmp := element.Value.Call("getAttribute", name)

		if !tmp.IsNull() {
			element.Attributes[name] = tmp.String()
			value = element.Attributes[name]
		}

	}

	return value

}

func (element *Element) SetAttribute(name string, value string) bool {

	var result bool = false

	check := validateXMLName(name)

	if check == nil {
		element.Attributes[name] = value
		element.Value.Call("setAttribute", name, value)
	} else {
		result = false
	}

	return result

}

func (element *Element) HasAttribute(name string) bool {

	var result bool = false

	check := validateXMLName(name)

	if check == nil {

		tmp := element.Value.Call("hasAttribute", name)

		if tmp.Truthy() {
			result = true
		}

	}

	return result

}

func (element *Element) QuerySelector(query string) *Element {

	var result *Element = nil

	value := element.Value.Call("querySelector", query)

	if !value.IsNull() && !value.IsUndefined() {
		child := ToElement(value)
		result = &child
	}

	return result

}

func (element *Element) QuerySelectorAll(query string) []*Element {

	var result []*Element

	values := element.Value.Call("querySelectorAll", query)

	for v := 0; v < values.Length(); v++ {

		value := values.Index(v)

		if !value.IsNull() && !value.IsUndefined() {

			child := ToElement(value)
			result = append(result, &child)

		}

	}

	return result

}

func (element *Element) SetClassName(value string) bool {

	var result bool = false

	value = strings.TrimSpace(value)

	element.Value.Set("className", value)
	element.ClassName = element.Value.Get("className").String()

	if element.ClassName == value {
		result = true
	}

	return result

}

func (element *Element) SetInnerHTML(value string) bool {

	element.Value.Set("innerHTML", value)
	element.InnerHTML = element.Value.Get("innerHTML").String()

	return true

}
