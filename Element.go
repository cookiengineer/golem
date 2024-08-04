package godom

import "godom/utils"
import "syscall/js"

type Element struct {
	id          string            `json:"id"`
	className   string            `json:"className"`
	attributes  map[string]string `json:"attributes"`
	textContent string            `json:"textContent"`
	innerHTML   string            `json:"innerHTML"`
	Value       js.Value          `json:"value"`
}

func (element *Element) Attributes() map[string]string {

	// TODO: Get Attributes again

	return element.attributes

}

func (element *Element) SetAttribute(name string, value string) error {

	var err error = nil

	err0 := utils.ValidateXMLName(name)

	if err0 == nil {
		element.attributes[name] = value
		element.Value.Call("setAttribute", name, value)
	} else {
		err = err0
	}

	return err

}

func (element *Element) HasAttributes() bool {

	var result bool = false

	value := element.Value.Call("hasAttribute")

	if value.Truthy() {
		result = true
	}

	return result

}
