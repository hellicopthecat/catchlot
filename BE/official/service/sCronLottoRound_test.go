package service

import (
	"context"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestSCronLottoRound(t *testing.T) {

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
		chromedp.Evaluate(`(() => {
  		const boxes = Array.from(document.querySelectorAll('.result-ballBox'));
  		const candidates = boxes.map(
				(box) => 
					Array.from(box.querySelectorAll('.result-ball'))
					.map((e)=>e.textContent.trim())).filter((balls)=>balls.length ===7);
			return candidates[candidates.length -1] || [];
			})()`, &texts),
	)
	if err != nil {
		t.Fatalf("스크래핑 실패: %v", err)
	}

	t.Logf("당첨번호: %v", texts)
	if len(texts) != 7 {
		t.Errorf("기대값 7개, 실제 %d개", len(texts))
	}
}
