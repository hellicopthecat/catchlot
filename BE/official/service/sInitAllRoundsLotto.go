package service

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/hellicopthecat/catchlot/official/request"
	"github.com/xuri/excelize/v2"
)

func (s *OfficialLottoService) SInitAllRoundsLotto(ctx context.Context) error {
	exists, err := s.repo.RExistsLottoRounds(ctx)
	if err != nil {
		log.Printf("❌ SInitAllRoundsLotto :: %s", err)
		return err
	}
	if exists {
		err := errors.New("이미 생성된 데이터들이 존재합니다.")
		log.Printf("❌ SInitAllRoundsLotto :: %s", err)
		return err
	}

	f, err := excelize.OpenFile("lotto.xlsx")
	if err != nil {
		log.Printf("❌ SInitAllRoundsLotto :: %s", err)
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("❌ SInitAllRoundsLotto :: %s", err)
		}
	}()

	rows, err := f.GetRows("Lotto")

	if err != nil {
		return err
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		var dto request.LottoRoundRequest
		uid, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		nums := make([]int, 8)
		for idx := 0; idx < 8; idx++ {
			n, err := strconv.Atoi(row[idx])
			if err != nil {
				CheckAtoi(err)
				return err
			}
			nums[idx] = n
		}

		dto.Id = uid.String()
		dto.Round_no = nums[0]
		dto.Draw_date = time.Now()
		dto.Bonus_number = nums[7]
		dto.Numbers = append(dto.Numbers, nums[1:7]...)

		err = s.repo.RCreateOfficialLotto(ctx, dto)
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckAtoi(err error) {
	log.Printf("❌ CheckAtoi Err :: %s", err)
}
