package premailer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicHTML(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
        h1, h2 {
        	color:red;
        }
        strong {
        	text-decoration:none
        }
        </style>
        </head>
        <body>
        <h1>Hi!</h1>
        <p><strong>Yes!</strong></p>
        </body>
        </html>`

	p := NewPremailerFromString(html)
	result_html, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, result_html, "<h1 style=\"color:red\">Hi!</h1>")
	assert.Contains(t, result_html, "<p><strong style=\"text-decoration:none\">Yes!</strong></p>")
}


func TestDataPremailerIgnore(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css" data-premailer="ignore">
        h1, h2 {
        	color:red;
        }
        strong {
        	text-decoration:none
        }
        </style>
        </head>
        <body>
        <h1>Hi!</h1>
        <p><strong>Yes!</strong></p>
        </body>
        </html>`

	p := NewPremailerFromString(html)
	result_html, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, result_html, "<h1>Hi!</h1>")
	assert.Contains(t, result_html, "<p><strong>Yes!</strong></p>")
}