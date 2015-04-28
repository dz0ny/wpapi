package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/contrib/newrelic"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/hoard"
)

func getTheme(name string) (string, bool) {
	url := fmt.Sprintf("https://wordpress.org/themes/%s", name)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", false
	}
	href, found := doc.Find("#themes .theme-actions a.button-primary").First().Attr("href")
	return href, found
}

func getPlugin(name string) (string, bool) {
	url := fmt.Sprintf("https://wordpress.org/plugins/%s", name)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", false
	}
	href, found := doc.Find("#plugin-description .button a").First().Attr("href")
	return href, found
}

func getThemeThumbnail(name string) (string, bool) {
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

func getThemeCached(name string) string {
	return hoard.Get(fmt.Sprintf("GetTheme_%s", name), func() (interface{}, *hoard.Expiration) {
		url, found := getTheme(name)
		if found {
			return url, hoard.Expires().AfterMinutes(5)
		} else {
			return "", hoard.Expires().AfterMinutes(30)
		}
	}).(string)
}

func getThemeThumbnailCached(name string) string {
	return hoard.Get(fmt.Sprintf("GetThemeThumbnail_%s", name), func() (interface{}, *hoard.Expiration) {
		url, found := getThemeThumbnail(name)
		if found {
			return url, hoard.Expires().AfterMinutes(5)
		} else {
			return "", hoard.Expires().AfterMinutes(30)
		}
	}).(string)
}

func getPluginCached(name string) string {
	return hoard.Get(fmt.Sprintf("GetPlugin_%s", name), func() (interface{}, *hoard.Expiration) {
		url, found := getPlugin(name)
		if found {
			return url, hoard.Expires().AfterMinutes(5)
		} else {
			return "", hoard.Expires().AfterMinutes(30)
		}
	}).(string)
}

func themeCached(c *gin.Context) {
	name := c.Params.ByName("name")
	url := getThemeCached(name)
	if url != "" {
		c.Redirect(http.StatusFound, url)
	} else {
		c.String(http.StatusNotFound, "Theme was not found")
	}
}

func themeCachedString(c *gin.Context) {
	name := c.Params.ByName("name")
	url := getThemeCached(name)
	if url != "" {
		c.String(http.StatusOK, url)
	} else {
		c.String(http.StatusNotFound, "Theme was not found")
	}
}

func pluginCached(c *gin.Context) {
	name := c.Params.ByName("name")
	url := getPluginCached(name)
	if url != "" {
		c.Redirect(http.StatusFound, url)
	} else {
		c.String(http.StatusNotFound, "Plugin was not found")
	}
}

func pluginCachedString(c *gin.Context) {
	name := c.Params.ByName("name")
	url := getPluginCached(name)
	if url != "" {
		c.String(http.StatusOK, url)
	} else {
		c.String(http.StatusNotFound, "Plugin was not found")
	}
}

func thumbCached(c *gin.Context) {
	name := c.Params.ByName("name")
	url := getThemeThumbnailCached(name)
	if url != "" {
		c.Redirect(http.StatusFound, url)
	} else {
		c.String(http.StatusNotFound, "Theme was not found")
	}
}

func getMainEngine() *gin.Engine {
	r := gin.Default()
	nr := os.Getenv("NEWRELIC")
	if nr != "" {
		r.Use(newrelic.NewRelic(nr, "wpapi", false))
	}
	r.GET("/theme/:name/zip", themeCachedString)
	r.GET("/plugin/:name/zip", pluginCachedString)

	r.HEAD("/theme/:name/download", themeCached)
	r.GET("/theme/:name/download", themeCached)
	r.HEAD("/plugin/:name/download", pluginCached)
	r.GET("/plugin/:name/download", pluginCached)
	r.GET("/theme/:name/thumbnail", thumbCached)
	return r
}

func main() {
	r := getMainEngine()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}
