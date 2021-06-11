package main
import (
	//"strings"
	//"fmt"
	//"net/url"
	"splashshopifymonitor/monitor"
	//"net/http"
	//"io/ioutil"
)

//var personalwebhook = "https://discord.com/api/webhooks/843003288487067710/rFw0ZGER5_s1QCOD_1h-496-DGpk7U1qATlkOD6rTajYAWhUCcv1BL81Naj-NVai8pQ6"
var r2rwebhook = "https://canary.discord.com/api/webhooks/852726307346841610/zhT7akZ2VnDLdrXuFFl8BdCDf0TanqRk0TSPskTm8fn-IhxvAHDsJI2cITGr2pJKLBWs"
 
func main() {
	a := monitor.Scraper{BaseURL: "https://splashstore2.myshopify.com/", Webhook: r2rwebhook,}

	a.Monitor()
	//postlocalhost()
}