# go-premailer

Inline styling for html in golang

# install
	
	go get github.com/vanng822/go-premailer/premailer

# Example

	import (
		"fmt"
		"github.com/vanng822/go-premailer/premailer"
		"log"
	)
	
	func main() {
		prem := premailer.NewPremailerFromFile(inputFile, premailer.NewOptions())
		html, err := prem.Transform()
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Println(html)
	}
	
# Commandline

	> go run main.go -i your_email.html
	> go run main.go -i your_mail.html -o process_mail.html
	
# Demo
	
http://premailer.isgoodness.com/
	
# Conversion endpoint

http://premailer.isgoodness.com/convert
	
	request POST:
		html(string)
	response:
		{result: output}
	