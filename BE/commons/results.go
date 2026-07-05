package commons

type Results struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   any    `json:"data"`
}
