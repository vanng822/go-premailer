package main

import (
	"flag"
	"fmt"
	"github.com/vanng822/go-premailer/premailer"
	"log"
	"os"
)

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

	prem := premailer.NewPremailerFromFile(inputFile)
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
