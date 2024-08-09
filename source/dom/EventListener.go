package dom

import "syscall/js"

var event_listener_id uint = 0

type EventListener struct {
	id       uint                  `json:"id"`
	callback EventListenerCallback `json:"callback"`
	scope    *js.Value             `json:"scope"`
}

type EventListenerCallback func(Event)

func ToEventListener(callback EventListenerCallback, scope *js.Value) EventListener {

	var listener EventListener

	listener.id = event_listener_id
	listener.callback = callback

	if scope != nil {
		listener.scope = scope
	} else {
		listener.scope = nil
	}

	event_listener_id += 1

	return listener

}
