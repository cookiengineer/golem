package dom

import "golem/types"
import "syscall/js"

type Event struct {
	Bubbles          bool       `json:"bubbles"`
	Cancelable       bool       `json:"cancelable"`
	Composed         bool       `json:"composed"`
	DefaultPrevented bool       `json:"defaultPrevented"`
	EventPhase       EventPhase `json:"eventPhase"`
	IsTrusted        bool       `json:"isTrusted"`
	Target           *Element   `json:"target"`
	// TimeStamp int `json:"timeStamp"`
	Type             EventType  `json:"type"`
	Value            *js.Value  `json:"value"`
}

func ToEvent(value js.Value) Event {

	var event Event

	event.Bubbles = types.ToBoolean(value.Get("bubbles"))
	event.Cancelable = types.ToBoolean(value.Get("cancelable"))
	event.Composed = types.ToBoolean(value.Get("composed"))
	event.DefaultPrevented = types.ToBoolean(value.Get("defaultPrevented"))
	event.EventPhase = EventPhase(types.ToInt(value.Get("eventPhase")))
	event.IsTrusted = types.ToBoolean(value.Get("isTrusted"))
	event.Type = EventType(types.ToString(value.Get("type")))

	target := ToElement(value.Get("target"))

	event.Target = &target
	event.Value = &value

	return event

}

func (event *Event) PreventDefault() {
	event.Value.Call("preventDefault")
}

func (event *Event) StopImmediatePropagation() {
	event.Value.Call("stopImmediatePropagation")
}

func (event *Event) StopPropagation() {
	event.Value.Call("stopPropagation")
}
