package dom

import "golem/types"
import "syscall/js"

type Event struct {
	Bubbles          bool       `json:"bubbles"`
	Cancelable       bool       `json:"cancelable"`
	Composed         bool       `json:"composed"`
	CurrentTarget    *Element   `json:"currentTarget"`
	DefaultPrevented bool       `json:"defaultPrevented"`
	EventPhase       EventPhase `json:"eventPhase"`
	IsTrusted        bool       `json:"isTrusted"`
	Target           *Element   `json:"target"`
	// TimeStamp int `json:"timeStamp"`
	Type             EventType  `json:"type"`
}

func ToEvent(value js.Value) Event {

	var event Event

	event.Bubbles = types.ToBoolean(value.Get("bubbles"))
	event.Cancelable = types.ToBoolean(value.Get("cancelable"))
	event.Composed = types.ToBoolean(value.Get("composed"))
	// TODO: CurrentTarget
	event.DefaultPrevented = types.ToBoolean(value.Get("defaultPrevented"))
	event.EventPhase = EventPhase(types.ToInt(value.Get("eventPhase")))
	event.IsTrusted = types.ToBoolean(value.Get("isTrusted"))
	// TODO: Target
	event.Type = EventType(types.ToString(value.Get("type")))

	return event

}

func (event *Event) PreventDefault() {

	// TODO

}

func (event *Event) StopImmedatePropagation() {

	// TODO

}

func (event *Event) StopPropagation() {

	// TODO

}
