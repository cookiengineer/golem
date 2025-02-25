package app

import "encoding/json"
import "gooey/dom"
import "gooey/storages"

type Main struct {
	Element *dom.Element     `json:"element"`
	Storage storages.Storage `json:"storage"`
	View    View             `json:"view"`
	Views   map[string]View  `json:"views"`
}

func (main *Main) Init(element *dom.Element) {

	main.Element = element
	main.Storage = storages.LocalStorage
	main.View    = nil
	main.Views   = make(map[string]View)

}

func (main *Main) SetView(name string, view View) {

	main.Views[name] = view

}

func (main *Main) ChangeView(name string) bool {

	var result bool = false

	view, ok := main.Views[name]

	if ok == true {

		if main.View != nil {
			main.View.Leave()
			main.View = nil
		}

		main.Element.SetAttribute("data-view", name)

		main.View = view
		main.View.Enter()

		result = true

	}

	return result

}

func (main *Main) ReadItem(name string, schema any) {

	buffer := main.Storage.GetItemBytes(name)

	if len(buffer) > 0 {
		json.Unmarshal(buffer, &schema)
	}

}

func (main *Main) SaveItem(name string, data any)  {

	buffer, err := json.Marshal(data)

	if err == nil {
		main.Storage.SetItem(name, buffer)
	}

}
