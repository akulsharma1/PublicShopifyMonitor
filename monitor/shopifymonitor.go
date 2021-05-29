package monitor

import (
	"time"
	"strings"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/aiomonitors/godiscord"
)
type Vars struct {
	Products []struct {
		ID          int64    `json:"id"`
		Title       string   `json:"title"`
		Handle      string   `json:"handle"`
		BodyHTML    string   `json:"body_html"`
		PublishedAt string   `json:"published_at"`
		CreatedAt   string   `json:"created_at"`
		UpdatedAt   string   `json:"updated_at"`
		Vendor      string   `json:"vendor"`
		ProductType string   `json:"product_type"`
		Tags        []string `json:"tags"`
		Variants    []struct {
			ID               int64       `json:"id"`
			Title            string      `json:"title"`
			Option1          string      `json:"option1"`
			Option2          interface{} `json:"option2"`
			Option3          interface{} `json:"option3"`
			Sku              string      `json:"sku"`
			RequiresShipping bool        `json:"requires_shipping"`
			Taxable          bool        `json:"taxable"`
			FeaturedImage    interface{} `json:"featured_image"`
			Available        bool        `json:"available"`
			Price            string      `json:"price"`
			Grams            int         `json:"grams"`
			CompareAtPrice   interface{} `json:"compare_at_price"`
			Position         int         `json:"position"`
			ProductID        int64       `json:"product_id"`
			CreatedAt        string      `json:"created_at"`
			UpdatedAt        string      `json:"updated_at"`
		} `json:"variants"`
		Images []struct {
			ID         int64         `json:"id"`
			CreatedAt  string        `json:"created_at"`
			Position   int           `json:"position"`
			UpdatedAt  string        `json:"updated_at"`
			ProductID  int64         `json:"product_id"`
			VariantIds []interface{} `json:"variant_ids"`
			Src        string        `json:"src"`
			Width      int           `json:"width"`
			Height     int           `json:"height"`
		} `json:"images"`
		Options []struct {
			Name     string   `json:"name"`
			Position int      `json:"position"`
			Values   []string `json:"values"`
		} `json:"options"`
	} `json:"products"`
}
type Scraper struct {
	BaseURL string

	ProxyList []string

	ProductTitle string
	ImageURL string
	Handle string
}
var webhook = "https://discord.com/api/webhooks/848075809594802197/HqGi437WTNC1a2ItlKbDZrJb0RJzA8_XryVWwtqMlU988xw_2ajt-jZCLdK_jZzb6uJ7"

func (t *Scraper)SendWebhook() {
	e := godiscord.NewEmbed(t.ProductTitle, "Website: "+t.BaseURL, t.BaseURL+"products/"+t.Handle)
	e.SetColor("#00FF00")
	e.SetThumbnail(t.ImageURL)
	e.SetFooter("Shopify Monitor • Made by splash#0003", "https://pbs.twimg.com/profile_images/1351753538066546690/bh72m_6R_400x400.png")
	e.SendToWebhook(webhook)
}
func VarRequest(url string) {
	req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Error initializing get page\n")
		}
		req.Header.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"90\", \"Google Chrome\";v=\"90\"")
		req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		req.Header.Set("Sec-Fetch-Site", "none")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Sec-Fetch-User", "?1")
		req.Header.Set("Sec-Fetch-Dest", "document")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("Error getting page, err %s\n", err)
		}
		defer resp.Body.Close()

		pageJson, _ := ioutil.ReadAll(resp.Body)
		if len(pageJson) <= 0 {
			fmt.Println("Failed parsing page, retrying...")
			ScrapeVars(url)
		} else {
			resp := pageJson
    		data := &Vars{}
    		_ = json.Unmarshal([]byte(resp), data)
			for i, value := range data.Products[0].Variants {
				VariantMap[i] = value.ID
				VariantString += fmt.Sprint(VariantMap[i])+"\n"
		
			}
		}
}
var VariantMap = make(map[int]int64)
var VariantString string = ""
func ScrapeVars(url string) {
	if !strings.Contains(url, "variant=") && url[len(url)-5:] == ".json"{
		fmt.Println(url[len(url)-5:])
		VarRequest(url)
		
	} else if strings.Contains(url, "variant=") && !(url[len(url)-5:] == ".json"){
		prodURL := url[:len(url)-23]+".json"
		VarRequest(prodURL)
		
	} else if strings.Contains(url, "variant=") && url[len(url)-5:] == ".json" {
		prodURL := url[:len(url)-28]+".json"
		VarRequest(prodURL)
		
	} else if !strings.Contains(url, "variant=") && !(url[len(url)-5:] == ".json") {
		url += ".json"
		fmt.Println(url)
		VarRequest(url)
		
	}
}

func (t *Scraper) Monitor() {
	loop:
		for {
			req, err := http.NewRequest("GET", t.BaseURL+"products.json?limit=25", nil)
			if err != nil {
				fmt.Println(err.Error())
				break loop
			}
			//fmt.Println(t.BaseURL+"products.json?limit=25")
			req.Header.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"90\", \"Google Chrome\";v=\"90\"")
			req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
			req.Header.Set("Upgrade-Insecure-Requests", "1")
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
			req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
			req.Header.Set("Sec-Fetch-Site", "none")
			req.Header.Set("Sec-Fetch-Mode", "navigate")
			req.Header.Set("Sec-Fetch-User", "?1")
			req.Header.Set("Sec-Fetch-Dest", "document")
			req.Header.Set("Accept-Language", "en-US,en;q=0.9")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println(err.Error())
				break loop
			}
			defer resp.Body.Close()
			pageJson, _ := ioutil.ReadAll(resp.Body)
			//bytePageHtml := []byte(pageJson)
			//fmt.Println(string(pageJson))
			switch resp.StatusCode {
			case 200:
				fmt.Println("[200] Monitoring")
			case 401:
				fmt.Printf("[401] Password Page Up for %v\n", t.BaseURL)
			case 429:
				fmt.Println("[429] Rate limited")
			case 500:
				fmt.Println("[500] Server error")
			}
			if len(pageJson) <= 0{
				fmt.Println("Failed parsing page")
			} else {
				//fmt.Println("in")
				resp := string(pageJson)
				data := &Vars{}
				_ = json.Unmarshal([]byte(resp), data)
				
				if data.Products[0].Title != t.ProductTitle{
					t.ProductTitle = data.Products[0].Title
					//if 
					//t.ImageURL = data.Products[0].Images[0].Src
					t.Handle = data.Products[0].Handle
					t.SendWebhook()
				}
			}
			time.Sleep(4000*time.Millisecond)
			
		}
	

}