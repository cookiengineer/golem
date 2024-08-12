package golem

import "golem/dom"
import "syscall/js"

var Window window

type window struct {
	listeners map[uint]dom.EventListener `json:"listeners"`
	Value     *js.Value                  `json:"value"`
}
