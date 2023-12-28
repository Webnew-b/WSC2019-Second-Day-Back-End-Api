package eventError

type EventSlugNotFoundError struct {
}

func (e *EventSlugNotFoundError) Error() string {
	return "Event not found"
}
