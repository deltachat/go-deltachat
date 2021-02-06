package deltachat

// #include <deltachat.h>
import "C"

type Event struct {
	event *C.dc_event_t
}

const EVENT_TYPES_DATA1_IS_INT = DC_EVENT_CHAT_MODIFIED |
	DC_EVENT_CONFIGURE_PROGRESS |
	DC_EVENT_CONTACTS_CHANGED |
	DC_EVENT_ERROR_NETWORK |
	DC_EVENT_IMEX_PROGRESS |
	DC_EVENT_INCOMING_MSG |
	DC_EVENT_LOCATION_CHANGED |
	DC_EVENT_MSG_DELIVERED |
	DC_EVENT_MSG_FAILED |
	DC_EVENT_MSG_READ |
	DC_EVENT_MSGS_CHANGED |
	DC_EVENT_SECUREJOIN_INVITER_PROGRESS |
	DC_EVENT_SECUREJOIN_JOINER_PROGRESS

const EVENT_TYPES_DATA2_IS_INT = DC_EVENT_INCOMING_MSG |
	DC_EVENT_MSG_DELIVERED |
	DC_EVENT_MSG_FAILED |
	DC_EVENT_SECUREJOIN_INVITER_PROGRESS |
	DC_EVENT_SECUREJOIN_JOINER_PROGRESS |
	DC_EVENT_MSG_READ |
	DC_EVENT_MSGS_CHANGED

const EVENT_TYPES_ERROR = DC_EVENT_ERROR |
	DC_EVENT_ERROR_NETWORK |
	DC_EVENT_ERROR_SELF_NOT_IN_GROUP

const (
	DATA_TYPE_INT uint8 = iota
	DATA_TYPE_STRING
	DATA_TYPE_NIL
)

var dataTypeNames = map[uint8]string{
	DATA_TYPE_INT:    "int",
	DATA_TYPE_STRING: "string",
	DATA_TYPE_NIL:    "nil",
}

func (e *Event) GetId() int {
	return int(C.dc_event_get_id(e.event))
}

func (e *Event) GetData1Int() int {
	return int(C.dc_event_get_data1_int(e.event))
}

func (e *Event) GetData2String() *string {
	s := C.dc_event_get_data2_str(e.event)
	if s != nil {
		res := C.GoString(s)
		C.dc_str_unref(s)
		return &res
	} else {
		return nil
	}
}

func (e *Event) Free() {
	C.dc_event_unref(e.event)
}
