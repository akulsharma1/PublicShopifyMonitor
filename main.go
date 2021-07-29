package main
import (
	"fmt"
	"splashshopifymonitor/monitor"
	"splashshopifymonitor/taskengine"
)

var personalwebhook = "https://discord.com/api/webhooks/843003288487067710/rFw0ZGER5_s1QCOD_1h-496-DGpk7U1qATlkOD6rTajYAWhUCcv1BL81Naj-NVai8pQ6"
//var r2reewebhook = "https://discord.com/api/webhooks/852726307346841610/zhT7akZ2VnDLdrXuFFl8BdCDf0TanqRk0TSPskTm8fn-IhxvAHDsJI2cITGr2pJKLBWs"
//var spungemonitor = "https://discord.com/api/webhooks/857347729923244052/mgw9WEXzDynpPPpMINVmJbDRc92N_LD-xfNVry-aZ_PkoetnKEW-1prDGtli_5ufb3Ve"
//var personalwebhook2 = "https://discord.com/api/webhooks/852258573970440204/1MkAm2Hl0y6go-Szi1wsJODklBGWkXymjjqNxZc_PeHw2dN3_Sbb5umbnUnXryFUCqp3"
var splashxeewebhook = "https://discord.com/api/webhooks/865033735926054912/KMGivbgPaXgqhED3hKyXv4J1tVKuSD0-1S1SALM887RaDXs54WHH8Y1ciVfIScHQ9Rlc"
var splashxlfwebhook = "https://discord.com/api/webhooks/865034008776802305/M15PNyM11JhMQ6wYZ-SEbRsu_LqdaGaOn5Etio9O2eqntfmlO2P3gxrlQWuqbuTRBdGN"
var splashxtelfarwebhook = "https://discord.com/api/webhooks/865034095524315156/ztPlG6G1ypkpzwBN_by2iyWLwiCb05fAbunTgl6PozsmFan_m847XYfdBdU9F4cfyZab"
var r2reewebhook = "https://discord.com/api/webhooks/852726307346841610/zhT7akZ2VnDLdrXuFFl8BdCDf0TanqRk0TSPskTm8fn-IhxvAHDsJI2cITGr2pJKLBWs"
var r2rlfwebhook = "https://discord.com/api/webhooks/860336811636293643/7HYwfetYShNU2tp1KyZ5ZUTKq2mOqh417FfjuWQF62yB6mh5ttWEkjNZgFvapr7N8ndE"
var r2r2 = "https://discord.com/api/webhooks/867778373245337600/9NKTkYpIG4E2tybF_EAjR7Z6kYjC5o-ljbUzYxge2mQiT4D-xW39pmV0r-EMRiCG_R2t"

func main() {
	var website string
	var input string
	/* 
	fmt.Println("Welcome, reading website file")

	dat, err := ioutil.ReadFile("./websites/websites.txt")
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Read file. Ready to start monitoring for %s\n", string(dat)) */
	
	
	fmt.Println("Presets:\nEric Emanuel: 1\nLost Files: 2\nTelfar: 3\nCustom: 4")
	fmt.Println("Type the number which you want then press enter.")
	fmt.Scanln(&input)

	if input == "1" { // eric emanuel
		proxies, _ := taskengine.ReadProxies(true)
		a := monitor.Scraper{BaseURL: "https://ericemanuel.com/", Webhook: r2reewebhook, Proxies: proxies}
		a.Client, _ = a.NewClient()
		a.Monitor(splashxeewebhook)
	} else if input == "2" { // lost files
		proxies, _ := taskengine.ReadProxies(true)
		a := monitor.Scraper{BaseURL: "https://lostfiles.shop/", Webhook: r2rlfwebhook, Proxies: proxies}
		a.Client, _ = a.NewClient()
		a.Monitor(splashxlfwebhook)
	} else if input == "3" { // telfar
		proxies, _ := taskengine.ReadProxies(true)
		a := monitor.Scraper{BaseURL: "https://shop.telfar.net/", Webhook: personalwebhook, Proxies: proxies}
		a.Client, _ = a.NewClient()
		a.Monitor(splashxtelfarwebhook)
	} else if input == "4" { // other (r2r only)
		proxies, _ := taskengine.ReadProxies(true)
		fmt.Println("Enter discord webhook for custom option")
		var dwebhook string
		fmt.Scanln(&dwebhook)
		fmt.Println("What site do you want to begin monitoring for. Example: \nhttps://kith.com/")
		fmt.Scanln(&website)
		a := monitor.Scraper{BaseURL: website, Webhook: dwebhook, Proxies: proxies}
		a.Client, _ = a.NewClient()
		a.Monitor(personalwebhook)
	}
	
}