package main
import (
	//"strings"
	//"fmt"
	//"net/url"
	"splashshopifymonitor/monitor"
	//"net/http"
	//"io/ioutil"
)

var webhook = "https://discord.com/api/webhooks/843003288487067710/rFw0ZGER5_s1QCOD_1h-496-DGpk7U1qATlkOD6rTajYAWhUCcv1BL81Naj-NVai8pQ6"
 
func main() {
	a := monitor.Scraper{BaseURL: "https://splashstore2.myshopify.com/", Webhook: webhook,}

	a.Monitor()
	//postlocalhost()
}