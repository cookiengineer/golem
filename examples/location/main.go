package main

import "gooey"
import "gooey/console"
import "gooey/location"
import "gooey/timers"
import "time"

func main() {

	element := gooey.Document.QuerySelector("#location")

	tmp := location.Location.Href

	if tmp != "" {
		element.SetInnerHTML("This page is located at \"" + tmp + "\"!")
	}

	timers.SetTimeout(func() {
		console.Inspect(location.Location)
	}, 1000)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
