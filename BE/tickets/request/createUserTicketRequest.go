package request

type CreateUserTicketDto struct {
	Round_id int   `json:"round_id"`
	Rank     int   `json:"rank"`
	Numbers  []int `json:"numbers"`
}
type CreateUserTicketRequest struct {
	User_id     string
	Ticket_info []CreateUserTicketDto
}
