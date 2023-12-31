package api

type LogoutRequest struct {
	Token string `validate:"required"`
}

type LogoutRes struct {
	Message string `json:"message"`
}
