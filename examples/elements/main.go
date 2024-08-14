package main

import "golem"
import "golem/dom"
import "strconv"
import "time"

func main() {

	// divs := golem.Document.QuerySelectorAll("div")

	// for d := 0; d < len(divs); d++ {
	// 	fmt.Println(divs[d])
	// }

	// container := golem.Document.QuerySelector("#container")
	// li := container.QuerySelector("li")

	// li.SetAttribute("data-active", "yes")

	var count int = 0


	var listener = dom.ToEventListener(func(event dom.Event) {

		target := event.Target

		if target.Id == "clickable" {

			if target.ClassName == "active" {
				target.SetInnerHTML("Click me again! (" + strconv.Itoa(count) + ")")
				target.SetClassName("")
			} else {
				target.SetClassName("active")
				target.SetInnerHTML("Click me again! (" + strconv.Itoa(count) + ")")
				count++
			}

		}

	}, golem.Document.Value)

	golem.Document.AddEventListener("click", listener)

	for true {

		if count > 3 {
			golem.Document.RemoveEventListener("click", &listener)
			golem.Document.QuerySelector("#clickable").SetInnerHTML("Stop clicking me!")
			count = 0
		}

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
