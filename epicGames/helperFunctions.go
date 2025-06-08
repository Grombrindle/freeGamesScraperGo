package epicgames

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
)

func RandomDelay() chromedp.ActionFunc {
	return func(ctx context.Context) error {
		delay := time.Duration(rand.Intn(4000)+7000) * time.Millisecond
		return chromedp.Sleep(delay).Do(ctx)
	}
}

func RandomScreenshotName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("screenshots/screenshot_%d_%d.png", time.Now().Unix(), r.Intn(100000))
}

// func humanMoveWithRod(page *rod.Page, selector string) error {
//     el := page.MustElement(selector)
//     el.MustMoveMouseOut()
//     page.MustWaitRequestIdle()
//     time.Sleep(time.Duration(rand.Intn(3000)+1000) * time.Millisecond)
//     el.MustClick()
//     return nil
// }

//	func humanScroll(ctx context.Context, totalScroll, step int) error {
//		for current := 0; current < totalScroll; current += step {
//			scrollJS := `window.scrollTo(0, ` + fmt.Sprint(current) + `)`
//			if err := chromedp.Evaluate(scrollJS, nil).Do(ctx); err != nil {
//				return err
//			}
//			// Random delay between 100-300ms to simulate human pause
//			time.Sleep(time.Duration(rand.Intn(200)+100) * time.Millisecond)
//		}
//		return nil
//	}
