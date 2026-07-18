package service

import "github.com/hellicopthecat/catchlot/official/repo"

type OfficialLottoService struct {
	repo *repo.OfficialLottoRepo
}

func InitOfficialLottoService(repo *repo.OfficialLottoRepo) *OfficialLottoService {
	return &OfficialLottoService{
		repo: repo,
	}
}
