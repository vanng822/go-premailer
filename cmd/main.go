package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vanng822/go-premailer/premailer"
)

func main() {
	var (
		inputFile           string
		outputFile          string
		removeClasses       bool
		skipCssToAttributes bool
		text                bool
	)
	flag.StringVar(&inputFile, "i", "", "Input file")
	flag.StringVar(&outputFile, "o", "", "Output file")
	flag.BoolVar(&text, "text", false, "Output only text")
	flag.BoolVar(&removeClasses, "remove-classes", false, "Remove class attribute")
	flag.BoolVar(&skipCssToAttributes, "skip-css-to-attributes", false, "No copy of css property to html attribute")
	flag.Parse()
	if inputFile == "" {
		flag.Usage()
		return
	}
	options := premailer.NewOptions()
	options.RemoveClasses = removeClasses
	options.CssToAttributes = !skipCssToAttributes
	prem, err := premailer.NewPremailerFromFile(inputFile, options)
	if err != nil {
		log.Fatal(err)
	}
	html, err := prem.Transform()
	if err != nil {
		log.Fatal(err)
	}
	txt, err := prem.TransformText()
	if err != nil {
		log.Fatal(err)
	}
	if outputFile != "" {
		fd, err := os.Create(outputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()
		if text {
			_, err = fd.WriteString(txt)
		} else {
			_, err = fd.WriteString(html)
		}
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if text {
		fmt.Println(txt)
	} else {
		fmt.Println(html)
	}

}
