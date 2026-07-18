package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hellicopthecat/catchlot/official/request"
)

func (s *OfficialLottoService) SCreateOfficialLotto(ctx context.Context, dtos []request.LottoRoundRequest) error {
	for _, dto := range dtos {

		round_id, err := s.repo.RFindOfficialLottoByRound(ctx, dto.Round_no)
		if err != nil {
			return err
		}
		if round_id != nil {
			return errors.New("이미 등록된 회차입니다.")
		}

		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		dto.Id = id.String()

		err = s.repo.RCreateOfficialLotto(ctx, dto)
		if err != nil {
			return err
		}
	}
	return nil
}
