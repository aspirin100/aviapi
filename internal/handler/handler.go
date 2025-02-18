package handler

type TicketManager interface{
	GetTicketList()
	EditTicket()
	RemoveTicketInfo()
}

type Handler struct{

}

func New() *Handler {
	return &Handler{}
}