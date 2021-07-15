package main
import (
	"fmt"
	"splashshopifymonitor/monitor"
	"splashshopifymonitor/taskengine"
)

var personalwebhook = "https://discord.com/api/webhooks/843003288487067710/rFw0ZGER5_s1QCOD_1h-496-DGpk7U1qATlkOD6rTajYAWhUCcv1BL81Naj-NVai8pQ6"
//var r2reewebhook = "https://discord.com/api/webhooks/852726307346841610/zhT7akZ2VnDLdrXuFFl8BdCDf0TanqRk0TSPskTm8fn-IhxvAHDsJI2cITGr2pJKLBWs"
var spungemonitor = "https://discord.com/api/webhooks/857347729923244052/mgw9WEXzDynpPPpMINVmJbDRc92N_LD-xfNVry-aZ_PkoetnKEW-1prDGtli_5ufb3Ve"
//var personalwebhook2 = "https://discord.com/api/webhooks/852258573970440204/1MkAm2Hl0y6go-Szi1wsJODklBGWkXymjjqNxZc_PeHw2dN3_Sbb5umbnUnXryFUCqp3"
var splashxeewebhook = "https://discord.com/api/webhooks/865033735926054912/KMGivbgPaXgqhED3hKyXv4J1tVKuSD0-1S1SALM887RaDXs54WHH8Y1ciVfIScHQ9Rlc"
var splashxlfwebhook = "https://discord.com/api/webhooks/865033735926054912/KMGivbgPaXgqhED3hKyXv4J1tVKuSD0-1S1SALM887RaDXs54WHH8Y1ciVfIScHQ9Rlc"
var splashxtelfarwebhook = "https://discord.com/api/webhooks/865033735926054912/KMGivbgPaXgqhED3hKyXv4J1tVKuSD0-1S1SALM887RaDXs54WHH8Y1ciVfIScHQ9Rlc"
var r2reewebhook = "https://discord.com/api/webhooks/852726307346841610/zhT7akZ2VnDLdrXuFFl8BdCDf0TanqRk0TSPskTm8fn-IhxvAHDsJI2cITGr2pJKLBWs"
var r2rlfwebhook = "https://discord.com/api/webhooks/860336811636293643/7HYwfetYShNU2tp1KyZ5ZUTKq2mOqh417FfjuWQF62yB6mh5ttWEkjNZgFvapr7N8ndE"

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
	
	fmt.Println("Welcome, what site do you want to begin monitoring for. Example:")
	fmt.Println("https://kith.com/")
	fmt.Scanln(&website)
	fmt.Println("Type 1 then press enter to begin monitoring.")
	fmt.Scanln(&input)

	if input == "1" { // eric emanuel
		proxies, _ := taskengine.ReadProxies(true)
		a := monitor.Scraper{BaseURL: website, Webhook: r2reewebhook, Proxies: proxies}
		a.Client, _ = a.NewClient()
		a.Monitor(splashxeewebhook)
	} else if input == "2" { // lost files
		proxies, _ := taskengine.ReadProxies(true)
		a := monitor.Scraper{BaseURL: website, Webhook: r2rlfwebhook, Proxies: proxies}
		a.Client, _ = a.NewClient()
		a.Monitor(splashxlfwebhook)
	} else if input == "3" { // telfar
		proxies, _ := taskengine.ReadProxies(true)
		a := monitor.Scraper{BaseURL: website, Webhook: personalwebhook, Proxies: proxies}
		a.Client, _ = a.NewClient()
		a.Monitor(splashxtelfarwebhook)
	} else if input == "4" { // other (r2r only)
		proxies, _ := taskengine.ReadProxies(true)
		a := monitor.Scraper{BaseURL: website, Webhook: personalwebhook, Proxies: proxies}
		a.Client, _ = a.NewClient()
		a.Monitor(personalwebhook)
	}
	
}