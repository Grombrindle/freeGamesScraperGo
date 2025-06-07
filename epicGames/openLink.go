package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	//    h "github.com/go-rod/rod"
	// "github.com/go-rod/rod/lib/input"
	// "github.com/teocci/go-chrome-cookies/chrome"
	// un "github.com/Davincible/chromedp-undetected"
	// cu "github.com/lrakai/chromedp-undetected"
	// "github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

// Add random delays between actions
func randomDelay() chromedp.ActionFunc {
	return func(ctx context.Context) error {
		delay := time.Duration(rand.Intn(5000)+8000) * time.Millisecond // 1-4 seconds
		return chromedp.Sleep(delay).Do(ctx)
	}
}

// func humanMoveWithRod(page *rod.Page, selector string) error {
//     el := page.MustElement(selector)
//     el.MustMoveMouseOut()
//     page.MustWaitRequestIdle()
//     time.Sleep(time.Duration(rand.Intn(3000)+1000) * time.Millisecond)
//     el.MustClick()
//     return nil
// }

// func humanScroll(ctx context.Context, totalScroll, step int) error {
// 	for current := 0; current < totalScroll; current += step {
// 		scrollJS := `window.scrollTo(0, ` + fmt.Sprint(current) + `)`
// 		if err := chromedp.Evaluate(scrollJS, nil).Do(ctx); err != nil {
// 			return err
// 		}
// 		// Random delay between 100-300ms to simulate human pause
// 		time.Sleep(time.Duration(rand.Intn(200)+100) * time.Millisecond)
// 	}
// 	return nil
// }

func main() {

	chromeProfilePath := filepath.Join(
		os.Getenv("LOCALAPPDATA"),
		"Google",
		"Chrome",
		"User Data",
		"Default",
		// "Network",
		// "Cookies",
	)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserDataDir(chromeProfilePath),
		chromedp.WindowSize(1280, 800),
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("profile-directory", "Default"),
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("disable-infobars", true),
		chromedp.Flag("disable-notifications", true),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("excludeSwitches", "enable-automation"),
		chromedp.Flag("useAutomationExtension", false),
		// chromedp.Flag("load-extension", `C:\Users\Damasco\AppData\Local\Google\Chrome\User Data\Default\Extensions\infdcenbdoibcacogknkjleclhnjdmfh\1.0.2_0`),
		// chromedp.Flag("disable-extensions-except", `C:\Users\Damasco\AppData\Local\Google\Chrome\User Data\Default\Extensions\infdcenbdoibcacogknkjleclhnjdmfh\1.0.2_0`),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)
	var links []string
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// _ = chromedp.Run(ctx,
	// 	chromedp.Navigate("https://example.com"),
	// 	chromedp.Sleep(3*time.Second),
	// )

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://store.epicgames.com/en-US/"),
		chromedp.Sleep(2*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 7; i++ {
		time.Sleep(300 * time.Millisecond)
		err := chromedp.Run(ctx,
			chromedp.Evaluate(`window.scrollBy(0, 350)`, nil),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	time.Sleep(5000 * time.Millisecond)
	err = chromedp.Run(ctx,
		chromedp.Evaluate(`Array.from(document.querySelectorAll('div[data-component="VaultOfferCard"] a')).map(a => a.href)`, &links),
	)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Free games found:")

	// time.Sleep(40000 * time.Millisecond)
	// os.Exit(0)

	// var nodes []*chromedp.Node
	// if err := humanScroll(ctx, 500, 100); err != nil {
	// 	log.Fatal(err)
	// }
	for i, link := range links {
		fmt.Printf("%d: %s\n", i+1, link)
		err := chromedp.Run(ctx,
			chromedp.Navigate(link),
			chromedp.Sleep(2*time.Second),
		)
		if err != nil {
			// Element is visible
			log.Println("  what ? ")
		}
		timeoutCtx, cancel := context.WithTimeout(ctx, 7*time.Second) // 5-second timeout
		defer cancel()
		err = chromedp.Run(timeoutCtx,
			chromedp.WaitVisible(`//button[.//span[text()="In Library"]]`, chromedp.BySearch),
		)

		if err == nil {
			// Element is visible
			fmt.Println("  Element is visible, do something  ")
			continue
		} else {
			// Element not visible or not found within timeout
			fmt.Println("  Element not visible or timeout:   ", err)
		}

		err = chromedp.Run(ctx,
			randomDelay(),

			// cu.LoadCookiesFromFile("C:/Users/Damasco/Documents/test/store.epicgames.com.cookies.json"),
			// chromedp.Navigate("https://store.epicgames.com/en-US/"),

			// humanMove(),
			// chromedp.WaitVisible(`div[aria-label="Free Games"]`),

			// Wait for the game cards to load
			// chromedp.Nodes(`div[data-component="VaultOfferCard"] a`, &nodes, chromedp.ByQueryAll),
			// chromedp.Sleep(6*time.Second),

			// Click on the first free game (DEATHLOOP)

			// chromedp.Click(`#app-main-content > div.css-1dnikhe > div > div > div > div:nth-child(5) > div:nth-child(2) > span:nth-child(1) > div > div > section > div > div:nth-child(1) > div > div > a`, chromedp.NodeVisible),
			// randomDelay(),

			// chromedp.Click(`/html/body/div[1]/div/div/div[4]/main/div[2]/div/div/div/div[3]/div[2]/span[1]/div/div/section/div/div[2]/div/div/a/div/div[1]/div[1]/div/div/div/div/div`, chromedp.NodeVisible),
			// randomDelay(),
			// Wait for the game page to load
			// chromedp.WaitVisible(`h1`), // Wait for any h1 element to appear on the new page

			// chromedp.WaitVisible(`#app-main-content > div.css-1dnikhe > div > div > div > div.css-f0x796 > div > div > div > div > div:nth-child(2) > div > div.css-15s2kp8`, chromedp.ByQuery),
			// chromedp.Sleep(18*time.Second),
			chromedp.Click(`/html/body/div[1]/div/div/div[4]/main/div[2]/div/div/div/div[2]/div[4]/div/aside/div/div/div[5]/div[1]/button`, chromedp.NodeVisible, chromedp.BySearch),
			randomDelay(),

			chromedp.Click(`/html/body/div[1]/div/div[4]/div/div/div/div[2]/div[2]/div/button`, chromedp.NodeVisible, chromedp.BySearch),
			randomDelay(),

			chromedp.Click(`/html/body/div[1]/div[3]/div[2]/div/div[3]/button[2]`, chromedp.NodeVisible, chromedp.BySearch),
			randomDelay(),
			//IF WE WERE NOT LOGGED IN AND WANT TO DO AGE THINGYS BUT I SHOULDNOT NEED IT HERE (IT DOESNOT WORK BTW)
			/* chromedp.Click(`#month_toggle`, chromedp.ByID),
			chromedp.Sleep(1*time.Second),
			chromedp.Click(`//div[@class="css-2s7xvo" and @data-testid="title"]/span/span[text()="02"]`, chromedp.BySearch),*/
			// chromedp.Sleep(6*time.Second),
			// chromedp.WaitVisible(`#my-element`, chromedp.ByID),
		)

		if err != nil {
			log.Println("Get button not found or click failed, stopping loop.")
			continue
		}
		time.Sleep(5000 * time.Millisecond)
	}

	// var cookies []*network.Cookie
	// err = chromedp.Run(ctx,
	// 	chromedp.ActionFunc(func(ctx context.Context) error {
	// 		var err error
	// 		cookies, err = network.GetCookies().Do(ctx)
	// 		return err
	// 	}),
	// )
	// if err != nil {
	// 	log.Fatal("Failed to get cookies:", err)
	// }

	// log.Println("Retrieved cookies:")
	// for i, cookie := range cookies {
	// 	log.Printf("%d: %+v\n", i+1, cookie)
	// }
	time.Sleep(50000 * time.Millisecond)
	log.Println("Success! You should see your logged-in session.")
}
