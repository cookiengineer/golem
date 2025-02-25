package app

import "gooey/dom"

type Main struct {
	Element *dom.Element    `json:"element"`
	View    View            `json:"view"`
	Views   map[string]View `json:"views"`
}

func (main *Main) Init(element *dom.Element) {

	main.Element = element
	main.View    = nil
	main.Views   = make(map[string]View)

}

func (main *Main) SetView(name string, view View) {

	main.Views[name] = view

}

func (main *Main) ChangeView(name string, data any) bool {

	var result bool = false

	view, ok := main.Views[name]

	if ok == true {

		if main.View != nil {
			main.View.Leave(data)
			main.View = nil
		}

		main.Element.SetAttribute("data-view", name)

		main.View = view
		main.View.Enter(data)

		result = true

	}

	return result

}
