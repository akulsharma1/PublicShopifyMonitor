package main
import (
	"strings"
	"fmt"
	"net/url"
	"splashshopifymonitor/monitor"
	"net/http"
	"io/ioutil"
)
func main() {
	a := monitor.Scraper{BaseURL: "https://splashstore2.myshopify.com/"}

	a.Monitor()
	//postlocalhost()
}
func Postlocalhost() {
	//a, _ := url.Parse("https://localhost:8000")
	body := url.Values{}
	body.Add("name", "bob")
	body.Add("type", "fluffycat")
	b := strings.NewReader(body.Encode())
	fmt.Println(b)
	req, _ := http.NewRequest("POST", "http://localhost:8000", b)

	resp, err := http.DefaultClient.Do(req)
	resp.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	pageJson, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(pageJson))
	fmt.Println(resp.StatusCode)
} 