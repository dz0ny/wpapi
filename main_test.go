package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTheme(t *testing.T) {
	href, found := getTheme("editor")
	assert := assert.New(t)
	assert.True(found)
	assert.NotEmpty(href)
	assert.Contains(href, "https://")
}

func TestGetThemeThumbnail(t *testing.T) {
	href, found := getThemeThumbnail("editor")
	assert := assert.New(t)
	assert.True(found)
	assert.NotEmpty(href)
	assert.Contains(href, "https://")
}

func TestGetPlugin(t *testing.T) {
	href, found := getPlugin("akismet")
	assert := assert.New(t)
	assert.True(found)
	assert.NotEmpty(href)
	assert.Contains(href, "https://")
}

func TestGetThemeRequest(t *testing.T) {
	assert := assert.New(t)
	ts := GetMainEngine()
	req, _ := http.NewRequest("GET", "/theme/editor/download", nil)
	w := httptest.NewRecorder()

	ts.ServeHTTP(w, req)

	assert.Equal(w.Code, 301)
	assert.Contains(w.Header().Get("Location"), "https://")

}

func TestGetPluginRequestHead(t *testing.T) {
	assert := assert.New(t)
	ts := GetMainEngine()
	req, _ := http.NewRequest("HEAD", "/plugin/akismet/download", nil)
	w := httptest.NewRecorder()

	ts.ServeHTTP(w, req)

	assert.Equal(w.Code, 301)
	assert.Contains(w.Header().Get("Location"), "https://")

}

func TestGetThemeRequestHead(t *testing.T) {
	assert := assert.New(t)
	ts := GetMainEngine()
	req, _ := http.NewRequest("HEAD", "/theme/editor/download", nil)
	w := httptest.NewRecorder()

	ts.ServeHTTP(w, req)

	assert.Equal(w.Code, 301)
	assert.Contains(w.Header().Get("Location"), "https://")

}

func TestGetPluginRequest(t *testing.T) {
	assert := assert.New(t)
	ts := GetMainEngine()
	req, _ := http.NewRequest("GET", "/plugin/akismet/download", nil)
	w := httptest.NewRecorder()

	ts.ServeHTTP(w, req)

	assert.Equal(w.Code, 301)
	assert.Contains(w.Header().Get("Location"), "https://")

}

func TestGetThumbnailRequest(t *testing.T) {
	assert := assert.New(t)
	ts := GetMainEngine()
	req, _ := http.NewRequest("GET", "/theme/editor/thumbnail", nil)
	w := httptest.NewRecorder()

	ts.ServeHTTP(w, req)

	assert.Equal(w.Code, 301)
	assert.Contains(w.Header().Get("Location"), "https://")

}
