package main

import "golem"
import "golem/location"
import "time"

func main() {

	element := golem.Document.QuerySelector("#location")

	tmp := location.Location.Href

	if tmp != "" {
		element.SetInnerHTML("This page is located at \"" + tmp + "\"!")
	}

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
