package registrationsError

type AlreadyRegistrar struct {
}

func (e *AlreadyRegistrar) Error() string {
	return "User already registered"
}
