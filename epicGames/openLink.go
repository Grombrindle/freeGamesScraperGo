package epicgames

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	// "FreeGames/epicGames"
	"github.com/chromedp/chromedp"
	"github.com/gen2brain/beeep"
	//    h "github.com/go-rod/rod"
	// "github.com/go-rod/rod/lib/input"
	// "github.com/teocci/go-chrome-cookies/chrome"
	// un "github.com/Davincible/chromedp-undetected"
	// cu "github.com/lrakai/chromedp-undetected"
	// "github.com/chromedp/cdproto/cdp"
)

var (
	links []string
	res   []byte
)

func OpenLink() {

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
		chromedp.Flag("headless", true),
		chromedp.Flag("start-minimized", true),
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

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

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
		timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second) // 5-second timeout
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
			RandomDelay(),

			// cu.LoadCookiesFromFile("C:/Users/Damasco/Documents/test/store.epicgames.com.cookies.json"),
			// chromedp.Navigate("https://store.epicgames.com/en-US/"),

			// humanMove(),
			// chromedp.WaitVisible(`div[aria-label="Free Games"]`),

			// Wait for the game cards to load
			// chromedp.Nodes(`div[data-component="VaultOfferCard"] a`, &nodes, chromedp.ByQueryAll),
			// chromedp.Sleep(6*time.Second),

			// Click on the first free game (DEATHLOOP)

			// chromedp.Click(`#app-main-content > div.css-1dnikhe > div > div > div > div:nth-child(5) > div:nth-child(2) > span:nth-child(1) > div > div > section > div > div:nth-child(1) > div > div > a`, chromedp.NodeVisible),
			// RandomDelay(),

			// chromedp.Click(`/html/body/div[1]/div/div/div[4]/main/div[2]/div/div/div/div[3]/div[2]/span[1]/div/div/section/div/div[2]/div/div/a/div/div[1]/div[1]/div/div/div/div/div`, chromedp.NodeVisible),
			// RandomDelay(),
			// Wait for the game page to load
			// chromedp.WaitVisible(`h1`), // Wait for any h1 element to appear on the new page

			// chromedp.WaitVisible(`#app-main-content > div.css-1dnikhe > div > div > div > div.css-f0x796 > div > div > div > div > div:nth-child(2) > div > div.css-15s2kp8`, chromedp.ByQuery),
			// chromedp.Sleep(18*time.Second),
			// chromedp.Click(`/html/body/div[1]/div/div/div[4]/main/div[2]/div/div/div/div[2]/div[4]/div/aside/div/div/div[5]/div[1]/button`, chromedp.NodeVisible, chromedp.BySearch),
			chromedp.Click(`//button[.//span[text()="Get"]]`, chromedp.NodeVisible, chromedp.BySearch),
			RandomDelay(),

			chromedp.Click(`/html/body/div[1]/div/div[4]/div/div/div/div[2]/div[2]/div/button`, chromedp.NodeVisible, chromedp.BySearch),
			RandomDelay(),

			chromedp.Click(`/html/body/div[1]/div[3]/div[2]/div/div[3]/button[2]`, chromedp.NodeVisible, chromedp.BySearch),
			RandomDelay(),
			//IF WE WERE NOT LOGGED IN AND WANT TO DO AGE THINGYS BUT I SHOULDNOT NEED IT HERE (IT DOESNOT WORK BTW)
			/* chromedp.Click(`#month_toggle`, chromedp.ByID),
			chromedp.Sleep(1*time.Second),
			chromedp.Click(`//div[@class="css-2s7xvo" and @data-testid="title"]/span/span[text()="02"]`, chromedp.BySearch),*/
			// chromedp.Sleep(6*time.Second),
			// chromedp.WaitVisible(`#my-element`, chromedp.ByID),
		)
		if err != nil {
			log.Fatal(err)
		}
		err = beeep.Notify("My Go Program", "BRO COME SOLVE THE FUCKING CAPTCHA", "")
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(5000 * time.Millisecond)
		err = chromedp.Run(ctx,
			chromedp.CaptureScreenshot(&res),
		)
		if err != nil {
			log.Fatal(err)
		}
		err = chromedp.Run(ctx,
			chromedp.Evaluate(`window.scrollBy(0, 270)`, nil),
		)
		if err != nil {
			log.Fatal(err)
		}
		filename := RandomScreenshotName()
		err = os.WriteFile(filename, res, 0644)
		fmt.Printf("screen shot was taken and its now in : %s", filename)
		if err != nil {
			log.Fatal(err)
		}

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
	log.Println("Success!.")
}
