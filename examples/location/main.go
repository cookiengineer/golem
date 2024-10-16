package main

import "golem"
import "golem/console"
import "golem/location"
import "golem/timers"
import "time"

func main() {

	element := golem.Document.QuerySelector("#location")

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
