package deltachat

// #include <deltachat.h>
import "C"

// EventEmitter wraps dc_event_emitter_t
type EventEmitter struct {
	emitter *C.dc_event_emitter_t
}

func (e *EventEmitter) GetNextEvent() *Event {
        event := C.dc_get_next_event(e.emitter)
        return &Event {
                event: event,
        }
        // TODO unref
}

func (e *EventEmitter) Free() {
        C.dc_event_emitter_unref(e.emitter)
}
