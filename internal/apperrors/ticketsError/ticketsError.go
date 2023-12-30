package ticketsError

type NotAvailable struct{}

func (n NotAvailable) Error() string {
	return "Ticket is no longer available"
}
