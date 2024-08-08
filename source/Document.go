package golem

import "golem/dom"
import "syscall/js"

var Document document

func init() {

	Document = document{
		Value: js.Global().Get("document"),
	}

}

type document struct {
	Value js.Value `json:"value"`
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
