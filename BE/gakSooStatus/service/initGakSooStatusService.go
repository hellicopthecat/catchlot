package service

import "github.com/hellicopthecat/catchlot/gakSooStatus/repo"

type GakSooStatusService struct {
	repo *repo.GakSooStatusRepo
}

func InitGakSooStatusService(repo *repo.GakSooStatusRepo) *GakSooStatusService {
	return &GakSooStatusService{
		repo: repo,
	}
}
