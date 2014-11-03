package main

import (
	"testing"
)

func TestGetTheme(t *testing.T) {
	zip := GetTheme("editor")
	t.Log(zip)
	if zip == "" {
		t.Fail()
	}
}

func TestGetPLugin(t *testing.T) {
	zip := GetPlugin("akismet")
	t.Log(zip)
	if zip == "" {
		t.Fail()
	}
}
