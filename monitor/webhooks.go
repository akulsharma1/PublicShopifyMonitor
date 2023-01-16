package monitor

import (
	"fmt"
	"time"

	"github.com/akulsharma1/godiscord"
)
func (t *Scraper) SendNewProdWebhook() {
	currentTime := time.Now()
	e := godiscord.NewEmbed(t.ProductTitle, "", t.BaseURL+"products/"+t.Handle)
	e.SetAuthor(t.BaseURL, "", "")
	e.SetColor("#76a9db")
	e.SetThumbnail(t.ImageURL)
	e.AddField("Price", t.ProductPrice, false)
	for i := range VariantMaps {
		e.AddField("Size "+SizeMaps[i], "[ATC]"+ fmt.Sprintf("(%vcart/add?id=%v)", t.BaseURL, VariantMaps[i]), true)
	}
	e.Username = "splash"
	e.AddField("Quicktasks", fmt.Sprintf("[Wrath](http://localhost:32441/qt?input=%vproducts/%v) | ", t.BaseURL, t.Handle) + fmt.Sprintf("[Balko](http://localhost:6776/?url=%vproducts/%v) | ", t.BaseURL, t.Handle) + fmt.Sprintf("[Nebula](http://localhost:7392/quicktask?url=%vproducts/%v) | ", t.BaseURL, t.Handle) + fmt.Sprintf("[Cyber](https://cybersole.io/dashboard/tasks?quicktask=%vproducts/%v) | ", t.BaseURL, t.Handle) + fmt.Sprintf("[Prism](https://dashboard.prismaio.com?url=%vproducts/%v) | ", t.BaseURL, t.Handle) + fmt.Sprintf("[Sole](https://www.soleaio.dev/quick_tasks/create?taskGroup=Shopify&product=%vproducts/%v)", t.BaseURL, t.Handle), false)
	e.SetFooter("Shopify Monitor • " + currentTime.Format("[01/02 15:04:05]")+ " • splash#0003", "https://media.discordapp.net/attachments/766438340400250922/882795679532089354/splash.png")
	e.SendToWebhook(t.Webhook)
}

func (t *Scraper) SendProductsRemovedWebhook() {
	currentTime := time.Now()
	e := godiscord.NewEmbed("", fmt.Sprintf("All Products pulled on %v", t.BaseURL), "https://media.discordapp.net/attachments/766438340400250922/882795679532089354/splash.png")
	e.SetColor("#76a9db")
	e.SetAuthor(t.BaseURL, "", "")
	e.Username = "splash"
	e.SetFooter("Shopify Monitor • " + currentTime.Format("[01/02 15:04:05]")+ " • splash#0003", "https://media.discordapp.net/attachments/766438340400250922/882795679532089354/splash.png")
	e.SendToWebhook(t.Webhook)
}

func (t *Scraper) SendPwPageUpWebhook() {
	currentTime := time.Now()
	e := godiscord.NewEmbed("", fmt.Sprintf("Password page up on %v", t.BaseURL), "https://media.discordapp.net/attachments/766438340400250922/882795679532089354/splash.png")
	e.SetColor("#76a9db")
	e.SetAuthor(t.BaseURL, "", "")
	e.Username = "splash"
	e.SetFooter("Shopify Monitor • " + currentTime.Format("[01/02 15:04:05]")+ " • splash#0003", "https://media.discordapp.net/attachments/766438340400250922/882795679532089354/splash.png")
	e.SendToWebhook(t.Webhook)
}
func (t *Scraper) SendPwPageDownWebhook() {
	currentTime := time.Now()
	e := godiscord.NewEmbed("", fmt.Sprintf("Password page down on %v", t.BaseURL), "https://media.discordapp.net/attachments/766438340400250922/882795679532089354/splash.png")
	e.SetColor("#76a9db")
	e.SetAuthor(t.BaseURL, "", "")
	e.Username = "splash"
	e.SetFooter("Shopify Monitor • " + currentTime.Format("[01/02 15:04:05]")+ " • splash#0003", "https://media.discordapp.net/attachments/766438340400250922/882795679532089354/splash.png")
	e.SendToWebhook(t.Webhook)
}