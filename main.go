package main
import (
	//"strings"
	"fmt"
	//"net/url"
	"splashshopifymonitor/monitor"
	//"net/http"
	//"io/ioutil"
)

//var personalwebhook = "https://discord.com/api/webhooks/843003288487067710/rFw0ZGER5_s1QCOD_1h-496-DGpk7U1qATlkOD6rTajYAWhUCcv1BL81Naj-NVai8pQ6"
//var r2reewebhook = "https://canary.discord.com/api/webhooks/852726307346841610/zhT7akZ2VnDLdrXuFFl8BdCDf0TanqRk0TSPskTm8fn-IhxvAHDsJI2cITGr2pJKLBWs"
var r2rrichiewebhook = "https://discord.com/api/webhooks/853016234442489947/rIe1nUggovKQn81-uodKoJ7cXWjl3XBa8Nc5TcbLbb6vK8dmVjYr5R-yvrtmsv-YJ0LY"
 
func main() {
	var first string
	fmt.Println("Welcome, type 1 to start tasks")
	fmt.Scanln(&first)
	if first == "1" {
		a := monitor.Scraper{BaseURL: "https://richiele.com/", Webhook: r2rrichiewebhook,}
		a.Monitor()
	}
	

	
	//postlocalhost()
}