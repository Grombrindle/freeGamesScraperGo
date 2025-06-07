package main

import (
	// "fmt"
	vpn "freeGames/vpnOpen"
	"log"
	"strconv"
	"time"
	// "github.com/chromedp/chromedp"
	// "github.com/PuerkitoBio/goquery"
	// "github.com/gocolly/colly/v2"
)

var (
	pidNumber string
)

func main() {
	vpn.OpenVpn()
	time.Sleep(10 * time.Second)
	pidNumber = vpn.PID()

	pid, err := strconv.Atoi(pidNumber)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(20 * time.Second)

	vpn.KillProcess(pid)
}
