package golem

import "golem/dom"

type Screen struct {
	Width       uint               `json:"width"`
	Height      uint               `json:"height"`
	AvailWidth  uint               `json:"availWidth"`
	AvailHeight uint               `json:"availHeight"`
	ColorDepth  uint               `json:"colorDepth"`
	PixelDepth  uint               `json:"pixelDepth"`
	IsExtended  bool               `json:"isExtended"`
	Orientation *ScreenOrientation `json:"orientation"`
}

func (screen *Screen) AddEventListener(typ dom.EventType, listener dom.EventListener) bool {

	var result bool = false

	check := string(typ)

	if check == "change" || check == "orientationchange" {

		// TODO

	}

	return result

}

func (screen *Screen) RemoveEventListener(filter_type dom.EventType, listener *dom.EventListener) bool {

	var result bool = false

	// TODO

	return result

}

