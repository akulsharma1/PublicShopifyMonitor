package monitor

import (
	"math/rand"
	"net/url"
	"net/http"
	"splashshopifymonitor/client"
	"time"
	"fmt"
)

func (t *Scraper) NewClient() (*http.Client, error) {
	return &http.Client{}, nil
}
func (t *Scraper) SetupClient() {
	jar := client.NewExportableCookieJar()
	t.Client.Jar = jar
	t.CookieJar = *jar

	t.Client.Timeout = time.Millisecond * time.Duration(30000)
}
func newRoundTripper(u string) (*http.Transport, error) {
	proxyURL, err := url.Parse(u)

	if err != nil {
		return nil, err
	}

	return &http.Transport{Proxy: http.ProxyURL(proxyURL)}, nil
}
func (t *Scraper) RotateProxy() {
	var roundTripper http.RoundTripper
	var err error
	if len(t.Proxies) > 0 {
		proxyURL := t.Proxies[rand.Intn(len(t.Proxies))]
		roundTripper, err = newRoundTripper(proxyURL)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		roundTripper = &http.Transport{}
	}
	t.Client.Transport = roundTripper
}
