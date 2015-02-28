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

	p := NewPremailerFromString(html, nil)
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

	p := NewPremailerFromString(html, nil)
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

	p := NewPremailerFromString(html, nil)
	result_html, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, result_html, "<h1 style=\"color:red;width:100%\" width=\"100%\">Hi!</h1>")
	assert.Contains(t, result_html, "<p><strong style=\"text-decoration:none\">Yes!</strong></p>")
	assert.NotContains(t, result_html, "<style type=\"text/css\">")	
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

	p := NewPremailerFromString(html, nil)
	result_html, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, result_html, "<a href=\"/home\" style=\"color:green\">Yes!</a>")
	assert.Contains(t, result_html, "<style type=\"text/css\">")
}


func TestRemoveClass(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
        h1, h2 {
        	color:red;
        }
        .big {
        	font-size: 40px;
        }
        </style>
        </head>
        <body>
        <h1 class="big">Hi!</h1>
        <p><strong>Yes!</strong></p>
        </body>
        </html>`

	options := &Options{}
	options.RemoveClasses = true
	p := NewPremailerFromString(html, options)
	result_html, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, result_html, "<h1 style=\"color:red;font-size:40px\">Hi!</h1>")
	assert.Contains(t, result_html, "<p><strong>Yes!</strong></p>")
}


func TestCssToAttributesFalse(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
        h1, h2 {
        	color:red;
        }
        .wide {
        	width: 1000px;
        }
        </style>
        </head>
        <body>
        <h1 class="wide">Hi!</h1>
        <p><strong>Yes!</strong></p>
        </body>
        </html>`

	options := &Options{}
	options.CssToAttributes = false
	p := NewPremailerFromString(html, options)
	result_html, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, result_html, "<h1 class=\"wide\" style=\"color:red;width:1000px\">Hi!</h1>")
	assert.Contains(t, result_html, "<p><strong>Yes!</strong></p>")
}