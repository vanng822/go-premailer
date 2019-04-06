package premailer

import (
	"bytes"
	"testing"
)

var (
	testBenmarkBuf       *bytes.Buffer
	testBenmarkHtmlBytes []byte
	testBenmarkHtml      string
)

func init() {
	testBenmarkBuf = new(bytes.Buffer)
	testString := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
        h1 {
        	width: 50px;
        	color:red;
        }
        h2 {
        	vertical-align: top;
        }
        h3 {
		    	text-align: right;
				}
        strong {
        	text-decoration:none
        }
        div {
        	background-color: green
        }
        </style>
        </head>
        <body>
        <h1>Hi!</h1>
        <h2>There</h2>
        <h3>Hello</h3>
        <p><strong>Yes!</strong></p>
        <div>Green color</div>
        </body>
        </html>`
	testBenmarkBuf.WriteString(testString)
	testBenmarkHtml = testString
	testBenmarkHtmlBytes = []byte(testString)
}

func BenchmarkPremailerBasicHTMLBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		p, _ := NewPremailerFromBytes(testBenmarkHtmlBytes, nil)
		p.Transform()
	}
}

func BenchmarkPremailerBasicHTML(b *testing.B) {
	for n := 0; n < b.N; n++ {
		p, _ := NewPremailerFromString(testBenmarkHtml, nil)
		p.Transform()
	}
}

func BenchmarkPremailerBasicHTMLBytes2String(b *testing.B) {
	for n := 0; n < b.N; n++ {
		p, _ := NewPremailerFromString(testBenmarkBuf.String(), nil)
		p.Transform()
	}
}
