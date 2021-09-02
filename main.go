package main
import (
	"fmt"
	"splashshopifymonitor/monitor"
	"splashshopifymonitor/taskengine"
)

func main() {
	var website string
	var input string
	var delay int
	
	fmt.Println("Presets:\nCustom: 1\nCustom with keywords: 2")
	fmt.Println("Type the number which you want then press enter.")
	fmt.Scanln(&input)

	if input == "1" { 
		proxies, _ := taskengine.ReadProxies(true)
		fmt.Println("Enter discord webhook for custom option")
		var dwebhook string
		fmt.Scanln(&dwebhook)
		fmt.Println("What site do you want to begin monitoring for. Example: \nhttps://kith.com/")
		fmt.Scanln(&website)
		fmt.Println("Enter delay (in milliseconds)")
		fmt.Scanln(&delay)
		a := monitor.Scraper{BaseURL: website, Webhook: dwebhook, Proxies: proxies}
		a.Client, _ = a.NewClient()
		a.Monitor(delay)
	} else if input == "2" {
		var kw string
		proxies, _ := taskengine.ReadProxies(true)
		fmt.Println("Enter discord webhook for custom option")
		var dwebhook string
		fmt.Scanln(&dwebhook)
		fmt.Println("What site do you want to begin monitoring for. Example: \nhttps://kith.com/")
		fmt.Scanln(&website)
		fmt.Println("Enter keywords. Example: \nyeezy,350,-dunk")
		fmt.Scanln(&kw)
		fmt.Println("Enter delay (in milliseconds)")
		fmt.Scanln(&delay)
		a := monitor.Scraper{BaseURL: website, Webhook: dwebhook, Proxies: proxies}
		a.Client, _ = a.NewClient()
		a.KwMonitor(kw, delay)
	}
	
}