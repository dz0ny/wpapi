package main

import (
	"testing"
)

func TestGetTheme(t *testing.T) {
	href, found := GetTheme("editor")
	t.Log(href)
	if href == "" || !found {
		t.Fail()
	}
}

func TestGetThemeThumbnail(t *testing.T) {
	href, found := GetThemeThumbnail("editor")
	t.Log(href)
	if href == "" || !found {
		t.Fail()
	}
}

func TestGetPLugin(t *testing.T) {
	href, found := GetPlugin("akismet")
	t.Log(href)
	if href == "" || !found {
		t.Fail()
	}
}
