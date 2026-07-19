package service

import (
	"context"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

type ResultTypess struct {
	Round   int
	Numbers []int
	Bonus   int
}

func TestSCronLottoRound(t *testing.T) {

	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var results ResultTypess

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
		t.Errorf("❌ SCronLottoRound :: %s", err)
		return
	}

	t.Logf("당첨번호: %v", results)
}
