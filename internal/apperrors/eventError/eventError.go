package eventError

const (
	EVENT_NOT_FOUND = "Event not found"
	INVALID_BODY    = "Invalid body"
	INVALID_PARAMS  = "Invalid params"
)

type EventSlugNotFoundError struct {
}

func (e *EventSlugNotFoundError) Error() string {
	return EVENT_NOT_FOUND
}
