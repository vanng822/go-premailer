package premailer

import (
	"github.com/PuerkitoBio/goquery"
	"os"
)

func NewPremailerFromFile(filename string) Premailer {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	d, err := goquery.NewDocumentFromReader(fd)
	if err != nil {
		panic(err)
	}
	return NewPremailer(d)
}
