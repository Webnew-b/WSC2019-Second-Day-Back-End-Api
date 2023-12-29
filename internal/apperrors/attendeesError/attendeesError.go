package attendeesError

const (
	ErrInvalidLoginMessage = "Invalid login"
	ErrInvalidTokenMessage = "Invalid Token"
)

type NotFound struct {
}

func (e *NotFound) Error() string {
	return "Invalid login"
}

type LoginKeyNotExist struct {
}

func (e *LoginKeyNotExist) Error() string {
	return "Invalid token"
}
