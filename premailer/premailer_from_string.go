package premailer

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func NewPremailerFromString(doc string) Premailer {
	read := strings.NewReader(doc)
	d, err := goquery.NewDocumentFromReader(read)
	if err != nil {
		panic(err)
	}
	return NewPremailer(d)
}
