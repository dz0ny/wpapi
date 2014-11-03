package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gorelic"
	"github.com/martini-contrib/throttle"
	"github.com/stretchr/hoard"
	"log"
	"os"
	"time"
)

var m *martini.Martini

func GetTheme(name string) string {
	url := fmt.Sprintf("https://wordpress.org/themes/%s", name)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	href, _ := doc.Find(".col-3 .button a").First().Attr("href")
	log.Println(href)
	return href
}

func GetThemeCached(name string) string {
	return hoard.Get(fmt.Sprintf("GetTheme_%s", name), func() (interface{}, *hoard.Expiration) {
		obj := GetTheme(name)
		return obj, hoard.Expires().AfterMinutes(30)
	}).(string)
}

func GetPlugin(name string) string {
	url := fmt.Sprintf("https://wordpress.org/plugins/%s", name)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	href, _ := doc.Find("#plugin-description .button a").First().Attr("href")
	log.Println(href)
	return href
}

func GetPluginCached(name string) string {
	return hoard.Get(fmt.Sprintf("GetPlugin_%s", name), func() (interface{}, *hoard.Expiration) {
		obj := GetTheme(name)
		return obj, hoard.Expires().AfterMinutes(30)
	}).(string)
}

func main() {
	m := martini.Classic()
	limits := throttle.Policy(&throttle.Quota{
		Limit:  60,
		Within: time.Second,
	})
	nr := os.Getenv("NEWRELIC")
	if nr != "" {
		m.Use(gorelic.Handler)
		gorelic.InitNewrelicAgent(nr, "wpapi", true)
	}

	m.Use(limits)
	m.Get("/theme/:name/zip", func(params martini.Params) string {
		return GetTheme(params["name"])
	})
	m.Get("/plugin/:name/zip", func(params martini.Params) string {
		return GetPlugin(params["name"])
	})

	m.Run()
}
