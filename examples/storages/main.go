package main

import "gooey"
import "gooey/storages"
import "gooey/timers"
import "strconv"
import "time"

func main() {

	element := gooey.Document.QuerySelector("#seconds")

	var seconds int = 0

	tmp := storages.SessionStorage.GetItemInt("seconds")

	if tmp > 0 {
		seconds = tmp
	}

	timers.SetInterval(func() {
		seconds++
		storages.SessionStorage.SetItem("seconds", seconds)
		element.SetInnerHTML("This page has been running for " + strconv.Itoa(seconds) + " seconds!")
	}, 1000)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
