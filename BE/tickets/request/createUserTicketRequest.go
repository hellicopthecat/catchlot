package request

type CreateUserTicketRequest struct {
	User_id  string
	Round_id int
	Rank     int
	Number   []int
}
