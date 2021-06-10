package monitor

import (
	"time"
	//"strings"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"reflect"
	//"github.com/stretchr/testify/require"
	//"github.com/emacampolo/gomparator"
	//"time"
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
	ProductPrice string
}

var VariantMaps = make(map[int]int64)
var SizeMaps = make(map[int]string)
var PreviousData = &Vars{}
func (t *Scraper) SetScraperEmpty() {
	//t.BaseURL = ""
	t.ProductTitle = ""
	t.ImageURL = ""
	t.Handle = ""
	t.ProductPrice = ""
	//VariantMaps = {}
	//SizeMaps = 
}
func (t *Scraper) Monitor() {
	sum := 0
	loop:
		for {
			req, err := http.NewRequest("GET", t.BaseURL+"products.json?limit=10", nil)
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
				
				resp := string(pageJson)
				data := &Vars{}
				_ = json.Unmarshal([]byte(resp), data)
				
				if sum == 0 {
					PreviousData = data
				} else {
					
					if len(data.Products) > len(PreviousData.Products) {
						//fmt.Println("in")
						difference := (len(data.Products) - len(PreviousData.Products))
						for i := 0; i <= difference-1; i++ {
							for j := range data.Products[i].Variants {
								VariantMaps[j] = data.Products[i].Variants[j].ID
								SizeMaps[j] = data.Products[i].Variants[j].Option1
							}
							t.ProductPrice = data.Products[i].Variants[0].Price
							t.ImageURL = data.Products[i].Images[0].Src
							t.ProductTitle = data.Products[i].Title
							t.Handle = data.Products[i].Handle
							t.SendNewProdWebhook()
							//t.SetScraperEmpty()
							fmt.Println("Sent product added webhook")
						}
					} else if len(data.Products) == 0 && len(PreviousData.Products) > 0 {
						//fmt.Println("in remove webhook")
						t.SendProductsRemovedWebhook()
						fmt.Println("Sent Products removed webhook")
					} else if len(data.Products) == len(PreviousData.Products) && len(data.Products) != 0 {
						for i := range data.Products {
							if reflect.DeepEqual(data.Products[i], PreviousData.Products[i]) {
								
							} else {
								prodExists := false
								for k := range data.Products {
									if reflect.DeepEqual(data.Products[i], PreviousData.Products[k]) {
										prodExists = true
										//break
									}
								}
								if !prodExists {
									for j := range data.Products[i].Variants {
										VariantMaps[j] = data.Products[i].Variants[j].ID
										SizeMaps[j] = data.Products[i].Variants[j].Option1
									}
									t.ProductPrice = data.Products[i].Variants[0].Price
									t.ImageURL = data.Products[i].Images[0].Src
									t.ProductTitle = data.Products[i].Title
									t.Handle = data.Products[i].Handle
									t.SendNewProdWebhook()
									//t.SetScraperEmpty()
									fmt.Println("Sent new product webhook")
								}
								//fmt.Println("found a difference")
								
							}
						}
					} else if len(data.Products) < len(PreviousData.Products) {

					}
				}
				PreviousData = data
			}
			sum++
			time.Sleep(4000*time.Millisecond)
			
		}
}