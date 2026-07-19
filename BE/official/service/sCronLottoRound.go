package service

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
	gakSooStatusRequest "github.com/hellicopthecat/catchlot/gakSooStatus/request"
	"github.com/hellicopthecat/catchlot/official/request"
	"github.com/robfig/cron/v3"
)

type ResultTypes struct {
	Round   int
	Numbers []int
	Bonus   int
}

func (s OfficialLottoService) SCronLottoRound(ctx context.Context) error {
	tl, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		log.Fatalf("❌ Load Location Failed :: %s", err)
	}
	c := cron.New(cron.WithLocation(tl))
	log.Println("✅ 크론 등록")
	_, err = c.AddFunc("46 20 * * 6", func() { // 크론 예약 시간 46 20 * * 6
		log.Println("✅ 크론 시작")
		result, err := fetchingNewRound()

		var newLottoRound request.LottoRoundRequest
		var newLottoRoundSlice []request.LottoRoundRequest

		uid, err := uuid.NewRandom()
		if err != nil {
			log.Printf("❌ SCronLottoRound uuid를 생성하지 못했습니다. :: %s", err)
			return
		}
		drawDate := time.Now().In(tl)

		newLottoRound.Id = uid.String()

		newLottoRound.Round_no = result.Round
		newLottoRound.Draw_date = drawDate
		newLottoRound.Numbers = result.Numbers
		newLottoRound.Bonus_number = result.Bonus
		newLottoRoundSlice = append(newLottoRoundSlice, newLottoRound)

		err = s.SCreateOfficialLotto(ctx, newLottoRoundSlice)
		if err != nil {
			log.Printf("❌ 크론 서비스 SCreateOfficialLotto 에러 :: %s", err)
			// return // 배포시 풀기
		}

		for _, nums := range newLottoRound.Numbers {
			var dto gakSooStatusRequest.GakSooStatusUpdateRequest
			dto.Bonus_count = false
			dto.Soo_id = strconv.Itoa(nums)
			if err = s.gakSooStatusRepo.RUpdateGakSooStatus(ctx, dto); err != nil {
				log.Printf("❌ Update GaskSooStatus :: %s", err)
				return
			}
		}
		log.Println("✅ 크론 일반수 상태 업데이트 끝.")

		var bonusDto gakSooStatusRequest.GakSooStatusUpdateRequest
		bonusDto.Bonus_count = true
		bonusDto.Soo_id = strconv.Itoa(newLottoRound.Bonus_number)
		if err = s.gakSooStatusRepo.RUpdateGakSooStatus(ctx, bonusDto); err != nil {
			log.Printf("❌ Update GaskSooStatus Bonus :: %s", err)
			return
		}
		log.Println("✅ 크론 보너스 상태 업데이트 끝.")
		log.Println("✅ 크론 작업 성공 끝.")
	})
	if err != nil {
		log.Printf("❌ 크론 등록 실패 :: %s", err)
	}
	c.Start()
	return nil
}

func fetchingNewRound() (*ResultTypes, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", true)) // 디버깅 시 false

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var results ResultTypes

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.dhlottery.co.kr/lt645/result"),
		chromedp.WaitVisible(".result-ballBox .result-ball", chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.Evaluate(`(() => {
  		const boxes = Array.from(document.querySelectorAll('.result-ballBox'));
  		const candidates = boxes.map(
				(box) => 
					Array.from(box.querySelectorAll('.result-ball'))
					.map((e)=>e.textContent.trim())).filter((balls)=>balls.length ===7);
			const balls = candidates[candidates.length -1] || [];

			const roundText = document.querySelectorAll('.result-txt > .ltEpsd');
			const t_round = roundText[roundText.length -1]?.textContent.trim() || '';
			const round = parseInt(t_round,10);

			return {
				round,
				numbers: balls.slice(0,6).map(Number),
				bonus: balls.length === 7 ? Number(balls[6]) : 0,				
			}
			})()`, &results),
	)
	if err != nil {
		log.Printf("❌ SCronLottoRound :: %s", err)
		return nil, err
	}

	log.Printf("당첨번호: %v", results)

	return &results, nil

}
