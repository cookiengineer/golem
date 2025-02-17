package elements

import "gooey/dom"

type Dialog struct {
	Element *dom.Element `json:"element"`
}

func ToDialog(value js.Value) Dialog {

	var dialog Dialog

	element := dom.ToElement(value)

	dialog.Element = &element

	return dialog

}

func (dialog *Dialog) Open() bool {

	var result bool = false

	if dialog.Element != nil {

		dialog.Element.Element.Value.Call("showModal")

		if dialog.Element.HasAttribute("open") {
			result = true
		}

	}

	return result

}

func (dialog *Dialog) Close() {

	var result bool = false

	if dialog.Element != nil {

		dialog.Element.Element.Value.Call("close")

		if dialog.Element.HasAttribute("open") == false {
			result = true
		}

	}

	return result

}
