package main

import "golem"
import "fmt"

func main() {

	divs := golem.Document.QuerySelectorAll("div")

	for d := 0; d < len(divs); d++ {
		fmt.Println(divs[d])
	}

	container := golem.Document.QuerySelector("#container")
	li := container.QuerySelector("li")

	li.SetAttribute("data-active", "yes")

}
