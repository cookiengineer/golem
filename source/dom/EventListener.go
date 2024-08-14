package dom

import "syscall/js"

var event_listener_id uint = 0

type EventListener struct {
	Id       uint                  `json:"id"`
	Callback EventListenerCallback `json:"callback"`
	Function *js.Func              `json:"function"`
	Scope    *js.Value             `json:"scope"`
}

type EventListenerCallback func(Event)

func ToEventListener(callback EventListenerCallback, scope *js.Value) EventListener {

	var listener EventListener

	listener.Id = event_listener_id
	listener.Callback = callback

	if scope != nil {
		listener.Scope = scope
	} else {
		listener.Scope = nil
	}

	event_listener_id += 1

	return listener

}
