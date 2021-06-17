package main
import (
	"fmt"
	"splashshopifymonitor/monitor"
	"splashshopifymonitor/taskengine"
)

//var personalwebhook = "https://discord.com/api/webhooks/843003288487067710/rFw0ZGER5_s1QCOD_1h-496-DGpk7U1qATlkOD6rTajYAWhUCcv1BL81Naj-NVai8pQ6"
//var r2reewebhook = "https://discord.com/api/webhooks/852726307346841610/zhT7akZ2VnDLdrXuFFl8BdCDf0TanqRk0TSPskTm8fn-IhxvAHDsJI2cITGr2pJKLBWs"
//var r2rrichiewebhook = "https://discord.com/api/webhooks/853016234442489947/rIe1nUggovKQn81-uodKoJ7cXWjl3XBa8Nc5TcbLbb6vK8dmVjYr5R-yvrtmsv-YJ0LY"
//var r2rbodegawebhook = "https://discord.com/api/webhooks/854535961601507338/kKACrV-FelmxK8VzHXjgf7Hb3dfeANXAHmXZ-8XT3wE4i4P0u72NkJH5OBpNGEA4sR-u"
var r2rlfwebhook = "https://discord.com/api/webhooks/853997135683125261/FGbFDvDfZIF8dcXt12VOWjArgZnxs14TKodIJu5sBCGzqJJlDLc3K0XlrUcwnaE83AWO"
//var r2r bladeewebhook = "https://discord.com/api/webhooks/855118313281159179/_5Knr86_tv6QbdsZ4FJ2LbRoVN-XipjphYGv1l4iIWbSJgct4eBHwppfL0m5B9SvckBm"
 
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

	if input == "1" {
		proxies, _ := taskengine.ReadProxies(true)
		a := monitor.Scraper{BaseURL: website, Webhook: r2rlfwebhook, Proxies: proxies}
		a.Client, _ = a.NewClient()
		a.Monitor()
	}
	
}