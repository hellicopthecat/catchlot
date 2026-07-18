package request

import "time"

type LottoRoundRequest struct {
	Id           string
	Round_no     int
	Draw_date    time.Time
	Bonus_number int
	Numbers      []int
}
