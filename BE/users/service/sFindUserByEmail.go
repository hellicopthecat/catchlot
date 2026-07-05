package service

func (s UserService) FindUserByEmail(email string) (string, error) {
	return s.repo.RFindUserByEmail(email)
}
