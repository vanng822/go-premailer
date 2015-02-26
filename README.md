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
		prem := premailer.NewPremailerFromFile(inputFile)
		html, err := prem.Transform()
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Println(html)
	}
	
# Commandline

	> go run main.go -i your_email.html
	> go run main.go -i your_mail.html -o process_mail.html
	