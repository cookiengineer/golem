package elements

import "gooey/dom"

type Dialog struct {
	Element *dom.Element `json:"element"`
}

func ToDialog(element *dom.Element) Dialog {

	var dialog Dialog

	dialog.Element = element

	return dialog

}

func (dialog *Dialog) Open() bool {

	var result bool = false

	if dialog.Element != nil {

		dialog.Element.SetAttribute("open", "")

		if dialog.Element.HasAttribute("open") {
			result = true
		}

	}

	return result

}

func (dialog *Dialog) Close() bool {

	var result bool = false

	if dialog.Element != nil {

		dialog.Element.RemoveAttribute("open")

		if dialog.Element.HasAttribute("open") == false {
			result = true
		}

	}

	return result

}
