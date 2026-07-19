package service

import (
	gakSooStatusRepo "github.com/hellicopthecat/catchlot/gakSooStatus/repo"
	"github.com/hellicopthecat/catchlot/official/repo"
)

type OfficialLottoService struct {
	repo             *repo.OfficialLottoRepo
	gakSooStatusRepo *gakSooStatusRepo.GakSooStatusRepo
}

func InitOfficialLottoService(repo *repo.OfficialLottoRepo, repo2 *gakSooStatusRepo.GakSooStatusRepo) *OfficialLottoService {
	return &OfficialLottoService{
		repo:             repo,
		gakSooStatusRepo: repo2,
	}
}
