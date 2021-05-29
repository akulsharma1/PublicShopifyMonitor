package main
import (
	"splashshopifymonitor/monitor"
)
func main() {
	a := monitor.Scraper{BaseURL: "https://splashstore2.myshopify.com/"}

	a.Monitor()
}