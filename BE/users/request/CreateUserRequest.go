package request

type CreateUserRequest struct {
	Id       string
	Email    string
	Social   string
	Nickname string
}

type UpdateUserRequest struct {
	Nickname     *string
	RefreshToken *string
	Email        string
}
