package app

import "gooey/dom"

type View interface {
	Init()
	GetElement(string) *dom.Element
	SetElement(string, *dom.Element)
	RemoveElement(string) bool
	Enter(any) bool
	Leave(any) bool
}

