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


func TestWithInline(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
        h1, h2 {
        	width: 50px;
        	color:red;
        }
        strong {
        	text-decoration:none
        }
        </style>
        </head>
        <body>
        <h1 style="width: 100%;">Hi!</h1>
        <p><strong>Yes!</strong></p>
        </body>
        </html>`

	p := NewPremailerFromString(html)
	result_html, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, result_html, "<h1 style=\"color:red;width:100%\" width=\"100%\">Hi!</h1>")
	assert.Contains(t, result_html, "<p><strong style=\"text-decoration:none\">Yes!</strong></p>")
}

func TestPseudoSelectors(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
        a:active {
        	color: red;
        	font-size: 12px;
        }
        a:first-child {
        	color: green;
        }
        </style>
        </head>
        <body>
        <h1>Hi!</h1>
        <p>
        	<a href="/home">Yes!</a>
        	<a href="/away">No!</a>
        </p>
        </body>
        </html>`

	p := NewPremailerFromString(html)
	result_html, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, result_html, "<a href=\"/home\" style=\"color:green\">Yes!</a>")
}
