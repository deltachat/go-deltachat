package deltachat

// #include <deltachat.h>
import "C"

import "runtime"

// EventEmitter wraps dc_event_emitter_t
type EventEmitter struct {
	emitter *C.dc_event_emitter_t
}

func (e *EventEmitter) GetNextEvent() *Event {
	cEvent := C.dc_get_next_event(e.emitter)
	event := &Event{
		event: cEvent,
	}
	runtime.SetFinalizer(event, (*Event).Free)
	return event
}

func (e *EventEmitter) Free() {
	C.dc_event_emitter_unref(e.emitter)
}
