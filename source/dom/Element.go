package dom

import "golem/utils"
import "syscall/js"
import "fmt"

type Element struct {
	id          string            `json:"id"`
	className   string            `json:"className"`
	attributes  map[string]string `json:"attributes"`
	textContent string            `json:"textContent"`
	innerHTML   string            `json:"innerHTML"`
	Value       *js.Value         `json:"value"`
}

func ToElement(value js.Value) Element {

	var element Element

	element.id = value.Get("id").String()
	element.className = value.Get("className").String()
	element.textContent = value.Get("textContent").String()
	element.innerHTML = value.Get("innerHTML").String()
	element.Value = &value

	element.attributes = make(map[string]string)
	element.Attributes()

	return element

}

func (element *Element) Attributes() map[string]string {

	attributes := element.Value.Get("attributes")

	for a := 0; a < attributes.Length(); a++ {

		name := attributes.Index(a).Get("name")
		value := attributes.Index(a).Get("value")

		fmt.Println(name, value)

	}

	// TODO: Get Attributes again via JS Syscall
	// TODO: Update element.attributes map

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
