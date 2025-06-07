package main

import (
	// "fmt"
	vpn "freeGames/vpnOpen"
	// "log"
	"time"
	// "github.com/chromedp/chromedp"
	// "github.com/PuerkitoBio/goquery"
	// "github.com/gocolly/colly/v2"
)

var (
	pidNumber   int
	appFinished bool = true
)

func main() {
	vpn.OpenVpn()
	time.Sleep(10 * time.Second)

	if appFinished {
		time.Sleep(20 * time.Second)
		pidNumber = vpn.PID()
		vpn.KillProcess(pidNumber)
	}
}
