package premailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkPremailerBasicHTMLBytes(b *testing.B) {
	html := []byte(`<html>
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
        </html>`)

	for n := 0; n < b.N; n++ {
		p, err := NewPremailerFromBytes(html, nil)
		assert.Nil(b, err)
		result_html, err := p.Transform()
		assert.NotNil(b, result_html)
		assert.Nil(b, err)
	}
}

func BenchmarkPremailerBasicHTML(b *testing.B) {
	html := `<html>
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

	for n := 0; n < b.N; n++ {
		p, err := NewPremailerFromString(html, nil)
		assert.Nil(b, err)
		result_html, err := p.Transform()
		assert.NotNil(b, result_html)
		assert.Nil(b, err)
	}
}

func BenchmarkPremailerBasicHTMLBytes2String(b *testing.B) {
	html := []byte(`<html>
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
        </html>`)

	for n := 0; n < b.N; n++ {
		p, err := NewPremailerFromString(string(html), nil)
		assert.Nil(b, err)
		result_html, err := p.Transform()
		assert.NotNil(b, result_html)
		assert.Nil(b, err)
	}
}
