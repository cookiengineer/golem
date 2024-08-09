package main

import "golem"
import "golem/dom"
import "fmt"

func main() {

	// divs := golem.Document.QuerySelectorAll("div")

	// for d := 0; d < len(divs); d++ {
	// 	fmt.Println(divs[d])
	// }

	// container := golem.Document.QuerySelector("#container")
	// li := container.QuerySelector("li")

	// li.SetAttribute("data-active", "yes")

	fmt.Println("main()")

	golem.Document.AddEventListener("DOMContentLoaded", dom.ToEventListener(func(event dom.Event) {

		fmt.Println(event.Target)

	}, golem.Document.Value), true)


}
