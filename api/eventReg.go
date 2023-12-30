package api

type EventRegRequestParams struct {
	OrgSlug string `validate:"required"`
	EvSlug  string `validate:"required"`
	Token   string `validate:"required"`
}

type EventRegRequestBody struct {
	TicketID   int64   `json:"ticket_id" validate:"required"`
	SessionIds []int64 `json:"session_ids" validate:"required|gt=0"`
}

type EventRegParams struct {
	*EventRegRequestParams
	*EventRegRequestBody
}

type EventRegRes struct {
	Message string `json:"message"`
}
