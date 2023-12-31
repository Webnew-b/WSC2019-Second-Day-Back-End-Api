package api

type LoginRequest struct {
	Lastname         string `json:"lastname" validate:"required"`
	RegistrationCode string `json:"registration_code" validate:"required"`
}

type LoginRes struct {
	Id        int64  `json:"-"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Token     string `json:"token" gorm:"-"`
}
