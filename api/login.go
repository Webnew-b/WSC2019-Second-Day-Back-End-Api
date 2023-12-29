package api

type LoginRequest struct {
	Lastname         string `json:"lastname" validate:"required"`
	RegistrationCode string `json:"registration_code" validate:"required"`
}

type LoginRes struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Token     string `json:"token" gorm:"-"`
}

type LogoutRequest struct {
	Token string `validate:"required"`
}

type LogoutRes struct {
	Message string `json:"message"`
}
