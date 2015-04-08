package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gorelic"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/throttle"
	"github.com/stretchr/hoard"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var m *martini.Martini

func GetTheme(name string) (string, bool) {
	url := fmt.Sprintf("https://wordpress.org/themes/%s", name)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", false
	}
	href, found := doc.Find("#themes .theme-actions a.button-primary").First().Attr("href")
	return href, found
}

func GetPlugin(name string) (string, bool) {
	url := fmt.Sprintf("https://wordpress.org/plugins/%s", name)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", false
	}
	href, found := doc.Find("#plugin-description .button a").First().Attr("href")
	return href, found
}

func GetThemeThumbnail(name string) (string, bool) {
	url := fmt.Sprintf("https://wordpress.org/themes/%s", name)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", false
	}
	href, found := doc.Find("#themes .screenshot img").First().Attr("src")
	if found {
		href = strings.Replace(href, "w=1142", "w=160", -1)
	}
	return href, found
}

func GetThemeCached(name string) string {
	return hoard.Get(fmt.Sprintf("GetTheme_%s", name), func() (interface{}, *hoard.Expiration) {
		url, found := GetTheme(name)
		if found {
			return url, hoard.Expires().AfterMinutes(5)
		} else {
			return "", hoard.Expires().AfterMinutes(30)
		}
	}).(string)
}

func GetThemeThumbnailCached(name string) string {
	return hoard.Get(fmt.Sprintf("GetThemeThumbnail_%s", name), func() (interface{}, *hoard.Expiration) {
		url, found := GetThemeThumbnail(name)
		if found {
			return url, hoard.Expires().AfterMinutes(5)
		} else {
			return "", hoard.Expires().AfterMinutes(30)
		}
	}).(string)
}

func GetPluginCached(name string) string {
	return hoard.Get(fmt.Sprintf("GetPlugin_%s", name), func() (interface{}, *hoard.Expiration) {
		url, found := GetPlugin(name)
		if found {
			return url, hoard.Expires().AfterMinutes(5)
		} else {
			return "", hoard.Expires().AfterMinutes(30)
		}
	}).(string)
}

func main() {
	m := martini.Classic()
	m.Use(gzip.All())
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
	m.Get("/theme/:name/zip", func(params martini.Params) (int, string) {
		url := GetThemeCached(params["name"])
		log.Println(url)
		if url != "" {
			return http.StatusOK, url
		} else {
			return http.StatusNotFound, "Theme was not found"
		}
	})
	m.Get("/plugin/:name/zip", func(params martini.Params) (int, string) {
		url := GetPluginCached(params["name"])
		log.Println(url)
		if url != "" {
			return http.StatusOK, url
		} else {
			return http.StatusNotFound, "Plugin was not found"
		}
	})

	m.Get("/theme/:name/download", func(r *http.Request, w http.ResponseWriter, params martini.Params) {
		url := GetThemeCached(params["name"])
		log.Println(url)
		if url != "" {
			http.Redirect(w, r, fmt.Sprintf("https:%s", url), 302)
			return
		} else {
			http.NotFound(w, r)
			return
		}
	})
	m.Get("/plugin/:name/download", func(r *http.Request, w http.ResponseWriter, params martini.Params) {
		url := GetPluginCached(params["name"])
		log.Println(url)
		if url != "" {
			http.Redirect(w, r, url, 302)
			return
		} else {
			http.NotFound(w, r)
			return
		}
	})

	m.Get("/theme/:name/thumbnail", func(r *http.Request, w http.ResponseWriter, params martini.Params) {
		url := GetThemeThumbnailCached(params["name"])
		if url != "" {
			http.Redirect(w, r, url, 301)
			return
		} else {
			http.NotFound(w, r)
			return
		}
	})

	m.Run()
}
