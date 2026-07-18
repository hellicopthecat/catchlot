package service

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

type LottoResult struct {
	ReturnValue    string `json:"returnValue"`
	DrwNo          int    `json:"drwNo"`
	DrwNoDate      string `json:"drwNoDate"`
	DrwtNo1        int    `json:"drwtNo1"`
	DrwtNo2        int    `json:"drwtNo2"`
	DrwtNo3        int    `json:"drwtNo3"`
	DrwtNo4        int    `json:"drwtNo4"`
	DrwtNo5        int    `json:"drwtNo5"`
	DrwtNo6        int    `json:"drwtNo6"`
	BnusNo         int    `json:"bnusNo"`
	FirstWinamnt   int64  `json:"firstWinamnt"`
	FirstPrzwnerCo int    `json:"firstPrzwnerCo"`
}

func (r OfficialLottoService) SCronLottoRound() error {

	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	var texts []string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.dhlottery.co.kr/lt645/result"),
		chromedp.WaitVisible(".result-ballBox .result-ball", chromedp.ByQuery),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('.result-ballBox .result-ball')).map(e => e.textContent.trim())`, &texts),
	)
	if err != nil {
		log.Printf("❌ SCronLottoRound :: %s", err)
		return err
	}

	log.Printf("당첨번호: %v", texts)

	return nil
}

// tl, err := time.LoadLocation("Asia/Seoul")
// if err != nil {
// 	log.Fatalf("❌ Load Location Failed :: %s", err)
// }
// c := cron.New(cron.WithLocation(tl))
// c.AddFunc("30 21 * * 6", fetchingNewRound)

// c.Start()

func fetchingNewRound() {}
