package main

import "golem"
import "golem/dom"
import "strconv"
import "time"

func main() {

	var count int = 0

	var listener = dom.ToEventListener(func(event dom.Event) {

		target := event.Target

		if target.Id == "clickable" {

			if target.ClassName == "active" {
				count++
				target.SetInnerHTML("Click me again! (" + strconv.Itoa(count) + ")")
				target.SetClassName("")
			} else {
				count++
				target.SetClassName("active")
				target.SetInnerHTML("Click me again! (" + strconv.Itoa(count) + ")")
			}

		}

	})

	golem.Document.AddEventListener("click", listener)

	for true {

		if count > 10 {

			golem.Document.RemoveEventListener("click", &listener)

			clickable := golem.Document.QuerySelector("#clickable")
			clickable.SetClassName("disabled")
			clickable.SetInnerHTML("Stop clicking me!")

			count = 0

		}

		// Do Nothing
		time.Sleep(100 * time.Millisecond)

	}

}
