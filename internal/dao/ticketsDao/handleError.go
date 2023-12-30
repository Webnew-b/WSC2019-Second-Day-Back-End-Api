package ticketsDao

import (
	"wscmakebygo.com/internal/apperrors/ticketsError"
)

func throwError() error {
	return &ticketsError.NotAvailable{}
}
