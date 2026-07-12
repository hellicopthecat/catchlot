package request

import "time"

type TFindUserTicketByIdRequest struct {
	Id int
}

type TFindUserTicketsRequest struct {
	Id     string
	Limit  int
	Offset int
}

type TFindUserTicketResponse struct {
	Id         int       `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`

	User_id     string `json:"user_id"`
	Round_id    int    `json:"round_id"`
	Rank        int    `json:"rank"`
	Bonus_match int    `json:"bonus_match"`
	Checked     int    `json:"checked"`

	Numbers []TPickNumberResponse `json:"numbers"`
}

type TPickNumberResponse struct {
	Pick_number    int    `json:"pick_number"`
	Pick_number_id string `json:"pick_number_id"`
}
