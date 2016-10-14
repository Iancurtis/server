package handlers

import (
	"html/template"
	"testing"
)

func TestTruncatedContent(t *testing.T) {
	p0 := Page{}
	p1 := Page{}
	p2 := Page{}
	p0.Content = "1234567890"
	p1.Content = "123"
	p2.Content = "12345678901234567890"
	if p0.TruncatedContent() != p0.Content[:10] {
		t.Error("trancated error p0:", p0.TruncatedContent(), p0.Content, p0.Content[:9])
	}
	if p1.TruncatedContent() != p1.Content {
		t.Error("trancated error p1")
	}
	c := template.HTML("1234567890...")
	if p2.TruncatedContent() != c {
		t.Error("trancated error p2:", p2.TruncatedContent(), c)
	}
}
