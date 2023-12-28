package organizerError

type OrganizerSlugNotFoundError struct {
}

func (e *OrganizerSlugNotFoundError) Error() string {
	return "Organizer not found"
}
