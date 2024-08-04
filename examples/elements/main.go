package main

import "godom"
import "fmt"

func main() {

	divs := godom.Document.QuerySelectorAll("div")

	for d := 0; d < len(divs); d++ {
		fmt.Println(divs[d])
	}

}
