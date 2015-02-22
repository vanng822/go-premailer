package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/vanng822/go-premailer/premailer"
	"log"
	"os"
)

func DocFromFile(filename string) (*goquery.Document, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	return goquery.NewDocumentFromReader(fd)
}

func main() {
	var (
		inputFile  string
		outputFile string
	)
	flag.StringVar(&inputFile, "i", "", "Input file")
	flag.StringVar(&outputFile, "o", "", "Output file")
	flag.Parse()
	if inputFile == "" {
		flag.Usage()
		return
	}
	doc, err := DocFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	prem := premailer.NewPremailer(doc)
	html, err := prem.Transform()
	if err != nil {
		log.Fatal(err)
	}
	if outputFile != "" {
		fd, err := os.Create(outputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()
		fd.WriteString(html)
	} else {
		fmt.Println(html)
	}
}
