package premailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPremailerBasicHTMLBytes(t *testing.T) {
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

	p, err := NewPremailerFromBytes(html, nil)
	assert.Nil(t, err)
	result_html, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, result_html, "<h1 style=\"width:50px;color:red\" width=\"50\">Hi!</h1>")
	assert.Contains(t, result_html, "<h2 style=\"vertical-align:top\" valign=\"top\">There</h2>")
	assert.Contains(t, result_html, "<h3 style=\"text-align:right\" align=\"right\">Hello</h3>")
	assert.Contains(t, result_html, "<p><strong style=\"text-decoration:none\">Yes!</strong></p>")
	assert.Contains(t, result_html, "<div style=\"background-color:green\" bgcolor=\"green\">Green color</div>")
}
