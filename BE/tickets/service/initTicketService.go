package service

import "github.com/hellicopthecat/catchlot/tickets/repo"

type TicketService struct {
	repo *repo.TicketRepo
}

func InitTicketService(repo *repo.TicketRepo) *TicketService {
	return &TicketService{
		repo: repo,
	}
}
