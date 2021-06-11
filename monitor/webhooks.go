package monitor

import (
	"time"
	"github.com/aiomonitors/godiscord"
	"fmt"
)
func (t *Scraper) SendNewProdWebhook() {
	currentTime := time.Now()
	e := godiscord.NewEmbed(t.ProductTitle, "", t.BaseURL+"products/"+t.Handle)
	e.SetAuthor(t.BaseURL, "", "")
	e.SetColor("#00f7f7")
	e.SetThumbnail(t.ImageURL)
	e.AddField("Price", t.ProductPrice, false)
	for i := range VariantMaps {
		e.AddField("Size "+SizeMaps[i], "[ATC]"+ fmt.Sprintf("(%vcart/add?id=%v)", t.BaseURL, VariantMaps[i]), true)
	}
	e.SetFooter("R2R • " + currentTime.Format("[01/02 15:04:05]")+ " • splash#0003", "https://media.discordapp.net/attachments/783026351531491361/852038253863632906/R2R.png?width=559&height=559")
	e.SendToWebhook(t.Webhook)
}

func (t *Scraper) SendProductsRemovedWebhook() {
	currentTime := time.Now()
	e := godiscord.NewEmbed("", fmt.Sprintf("All Products pulled on %v", t.BaseURL), "")
	e.SetColor("#76a9db")
	e.SetAuthor(t.BaseURL, "", "")
	e.SetFooter("R2R • " + currentTime.Format("[01/02 15:04:05]")+ " • splash#0003", "https://media.discordapp.net/attachments/783026351531491361/852038253863632906/R2R.png?width=559&height=559")
	e.SendToWebhook(t.Webhook)
}

func (t *Scraper) SendPwPageUpWebhook() {
	currentTime := time.Now()
	e := godiscord.NewEmbed("", fmt.Sprintf("Password page up on %v", t.BaseURL), "")
	e.SetColor("#76a9db")
	e.SetAuthor(t.BaseURL, "", "")
	e.SetFooter("R2R • " + currentTime.Format("[01/02 15:04:05]")+ " • splash#0003", "https://media.discordapp.net/attachments/783026351531491361/852038253863632906/R2R.png?width=559&height=559")
	e.SendToWebhook(t.Webhook)
}