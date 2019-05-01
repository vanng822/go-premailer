# go-premailer

Inline styling for HTML mail in golang

# Document
[![Build Status](https://travis-ci.org/vanng822/go-premailer.svg?branch=master)](https://travis-ci.org/vanng822/go-premailer)
[![GoDoc](https://godoc.org/github.com/vanng822/go-premailer/premailer?status.svg)](https://godoc.org/github.com/vanng822/go-premailer/premailer)
[![GoCover](http://gocover.io/_badge/github.com/vanng822/go-premailer/premailer)](http://gocover.io/github.com/vanng822/go-premailer/premailer)
[![Go Report Card](https://goreportcard.com/badge/github.com/vanng822/go-premailer)](https://goreportcard.com/report/github.com/vanng822/go-premailer)

# install

	go get github.com/vanng822/go-premailer/premailer

# Example

	import (
		"fmt"
		"github.com/vanng822/go-premailer/premailer"
		"log"
	)

	func main() {
		prem, err := premailer.NewPremailerFromFile(inputFile, premailer.NewOptions())
		if err != nil {
			log.Fatal(err)
		}
		
		html, err := prem.Transform()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(html)
	}

## Input

	<html>
	<head>
	<title>Title</title>
	<style type="text/css">
		h1 { width: 300px; color:red; }
		strong { text-decoration:none; }
	</style>
	</head>
	<body>
		<h1>Hi!</h1>
		<p><strong>Yes!</strong></p>
	</body>
	</html>

## Output

	<html>
	<head>
	<title>Title</title>
	</head>
	<body>
		<h1 style="color:red;width:300px" width="300">Hi!</h1>
		<p><strong style="text-decoration:none">Yes!</strong></p>
	</body>
	</html>



# Commandline

	> go run main.go -i your_email.html
	> go run main.go -i your_mail.html -o process_mail.html

# Demo

http://premailer.isgoodness.com/

# Conversion endpoint

http://premailer.isgoodness.com/convert

	request POST:
		html: your mail
		cssToAttributes: true|false
		removeClasses: true|false
	response:
		{result: output}
