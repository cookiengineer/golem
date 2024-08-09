package dom

import "golem/utils"
import "syscall/js"
import "fmt"

type Element struct {
	Id          string            `json:"id"`
	ClassName   string            `json:"className"`
	Attributes  map[string]string `json:"attributes"`
	TextContent string            `json:"textContent"`
	InnerHTML   string            `json:"innerHTML"`
	Value       *js.Value         `json:"value"`
}

func ToElement(value js.Value) Element {

	var element Element

	element.Id = value.Get("id").String()
	element.ClassName = value.Get("className").String()
	element.TextContent = value.Get("textContent").String()
	element.InnerHTML = value.Get("innerHTML").String()
	element.Value = &value

	element.Attributes = make(map[string]string)
	element.RefreshAttributes()

	return element

}

func (element *Element) RefreshAttributes() {

	attributes := element.Value.Get("attributes")

	for a := 0; a < attributes.Length(); a++ {

		name := attributes.Index(a).Get("name")
		value := attributes.Index(a).Get("value")

		fmt.Println(name, value)

	}

	// TODO: Get Attributes again via JS Syscall
	// TODO: Update element.attributes map

}

func (element *Element) GetAttribute(name string) string {

	var value string

	// TODO: Get attribute via JS Syscall

	return value

}

func (element *Element) SetAttribute(name string, value string) error {

	var err error = nil

	err0 := utils.ValidateXMLName(name)

	if err0 == nil {
		element.Attributes[name] = value
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
