package godom

import "syscall/js"
import "fmt"

var Document document

func init() {

	Document = document{
		Value: js.Global().Get("document"),
	}

}

type document struct {
	Value js.Value `json:"value"`
}

func (doc *document) QuerySelector(query string) Element {
}

func (doc *document) QuerySelectorAll(query string) []Element {

	var result []Element

	values := doc.Value.Call("querySelectorAll", query)

	for v := 0; v < values.Length(); v++ {

		value := values.Index(v)

		fmt.Println(value.Get("innerHTML").String())

	}

	return result

}
