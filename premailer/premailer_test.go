package premailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicHTML(t *testing.T) {
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

	p, err := NewPremailerFromString(html, nil)
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<h1 style=\"width:50px;color:red\" width=\"50\">Hi!</h1>")
	assert.Contains(t, resultHTML, "<h2 style=\"vertical-align:top\">There</h2>")
	assert.Contains(t, resultHTML, "<h3 style=\"text-align:right\">Hello</h3>")
	assert.Contains(t, resultHTML, "<p><strong style=\"text-decoration:none\">Yes!</strong></p>")
	assert.Contains(t, resultHTML, "<div style=\"background-color:green\">Green color</div>")
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

	p, err := NewPremailerFromString(html, nil)
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<h1>Hi!</h1>")
	assert.Contains(t, resultHTML, "<p><strong>Yes!</strong></p>")
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
        <h1 style="width: 100px;">Hi!</h1>
        <p><strong>Yes!</strong></p>
        </body>
        </html>`

	p, err := NewPremailerFromString(html, nil)
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<h1 style=\"color:red;width:100px\" width=\"100\">Hi!</h1>")
	assert.Contains(t, resultHTML, "<p><strong style=\"text-decoration:none\">Yes!</strong></p>")
	assert.NotContains(t, resultHTML, "<style type=\"text/css\">")
}

func TestWithZeroLength(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
        h1, h2 {
        	width: 0;
					height: 0;
        	color:red;
        }
        </style>
        </head>
        <body>
        <h1>Hi!</h1>
        <p><strong>Yes!</strong></p>
        </body>
        </html>`

	p, err := NewPremailerFromString(html, nil)
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<h1 style=\"width:0;height:0;color:red\" width=\"0\" height=\"0\">Hi!</h1>")
	assert.NotContains(t, resultHTML, "<style type=\"text/css\">")
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

	p, err := NewPremailerFromString(html, nil)
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<a href=\"/home\" style=\"color:green\">Yes!</a>")
	assert.Contains(t, resultHTML, "<style type=\"text/css\">")
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
        	width: 150px;
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
	p, err := NewPremailerFromString(html, options)
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<h1 style=\"color:red;font-size:40px;width:150px\">Hi!</h1>")
	assert.Contains(t, resultHTML, "<p><strong>Yes!</strong></p>")
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
	p, err := NewPremailerFromString(html, options)
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<h1 class=\"wide\" style=\"color:red;width:1000px\">Hi!</h1>")
	assert.Contains(t, resultHTML, "<p><strong>Yes!</strong></p>")
}

func TestWithImportant(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
        h1, h2 {
        	color:red;
        }
        p {
        	width: 100px !important;
        	color: blue
        }
        .wide {
        	width: 1000px;
        }
        </style>
        </head>
        <body>
        <h1>Hi!</h1>
        <p class="wide"><strong>Yes!</strong></p>
        </body>
        </html>`

	p, err := NewPremailerFromString(html, NewOptions())
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<h1 style=\"color:red\">Hi!</h1>")
	assert.Contains(t, resultHTML, "<p class=\"wide\" style=\"color:blue;width:100px\" width=\"100\"><strong>Yes!</strong></p>")
}

func TestWithKeepImportant(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
        h1, h2 {
        	color:red;
        }
        p {
        	width: 100px !important;
        	color: blue
        }
        .wide {
        	width: 1000px;
        }		
        </style>
        </head>
        <body>
        <h1>Hi!</h1>
        <p class="wide"><strong>Yes!</strong></p>
        </body>
        </html>`

	options := NewOptions()
	options.KeepBangImportant = true
	p, err := NewPremailerFromString(html, options)
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<h1 style=\"color:red\">Hi!</h1>")
	assert.Contains(t, resultHTML, "<p class=\"wide\" style=\"color:blue;width:100px !important\" width=\"100\"><strong>Yes!</strong></p>")

}

func TestWithMediaRule(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
        h1, h2 {
        	color:red;
        }
        p {
        	width: 100px !important;
        	color: blue
        }
        .wide {
        	width: 1000px;
        }
        @media all and (min-width: 62em) {
		    h1 {
		        font-size: 55px;
		        line-height: 60px;
		        padding-top: 0;
		        padding-bottom: 5px
		        }
		 }
        </style>
        </head>
        <body>
        <h1>Hi!</h1>
        <p class="wide"><strong>Yes!</strong></p>
        </body>
        </html>`

	p, err := NewPremailerFromString(html, NewOptions())
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<h1 style=\"color:red\">Hi!</h1>")
	assert.Contains(t, resultHTML, "<p class=\"wide\" style=\"color:blue;width:100px\" width=\"100\"><strong>Yes!</strong></p>")

	assert.Contains(t, resultHTML, "@media all and (min-width: 62em){")
	assert.Contains(t, resultHTML, "font-size: 55px !important;")
	assert.Contains(t, resultHTML, "line-height: 60px !important;")
	assert.Contains(t, resultHTML, "padding-bottom: 5px !important")
	assert.Contains(t, resultHTML, "padding-top: 0 !important")
}

func TestWithMediaAttribute(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
        h1, h2 {
        	color:red;
        }
        p {
        	width: 100px !important;
        	color: blue
        }
        .wide {
        	width: 1000px;
        }
       	</style>
      	<style type="text/css" media="all and (min-width: 62em)">
		    h1 {
		        font-size: 55px;
		        line-height: 60px;
		        padding-top: 0;
		        padding-bottom: 5px
		 }
        </style>
        <style type="text/css" media="all">
		h3 {
			color: black;
		}
		</style>
        <style>

        </style>
        </head>
        <body>
        <h1>Hi!</h1>
        <p class="wide"><strong>Yes!</strong></p>
        <h3>Hi, all media style!</h3>
        </body>
        </html>`

	p, err := NewPremailerFromString(html, NewOptions())
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<h1 style=\"color:red\">Hi!</h1>")
	assert.Contains(t, resultHTML, "<p class=\"wide\" style=\"color:blue;width:100px\" width=\"100\"><strong>Yes!</strong></p>")
	assert.Contains(t, resultHTML, "<h3 style=\"color:black\">Hi, all media style!</h3>")

	assert.Contains(t, resultHTML, "<style type=\"text/css\" media=\"all and (min-width: 62em)\">")
	assert.Contains(t, resultHTML, "font-size: 55px;")
	assert.Contains(t, resultHTML, "line-height: 60px;")
	assert.Contains(t, resultHTML, "padding-top: 0;")
	assert.Contains(t, resultHTML, "padding-bottom: 5px")

	assert.NotContains(t, resultHTML, "<style type=\"text/css\" media=\"all\">")
	assert.NotContains(t, resultHTML, "color: black;")
}

func TestIndexOutOfRange(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
        h1, h2 {
        	color:red;
        }
        p {
        	width: 100px !important;
        	color: blue
        }
        .wide {
        	width: 1000px;
        }
       	</style>
      	<style type="text/css" media="all and (min-width: 62em)">
		    h1 {
		        font-size: 55px;
		        line-height: 60px;
		        padding-top: 0;
		        padding-bottom: 5px
		 }
        </style>
        <style>
        	.some {
        		color: red;
        	}
        </style>
        </head>
        <body>
        <h1>Hi!</h1>
        <p class="wide"><strong>Yes!</strong></p>
        </body>
        </html>`

	p, err := NewPremailerFromString(html, NewOptions())
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<h1 style=\"color:red\">Hi!</h1>")
	assert.Contains(t, resultHTML, "<p class=\"wide\" style=\"color:blue;width:100px\" width=\"100\"><strong>Yes!</strong></p>")

	assert.Contains(t, resultHTML, "<style type=\"text/css\" media=\"all and (min-width: 62em)\">")
	assert.Contains(t, resultHTML, "font-size: 55px;")
	assert.Contains(t, resultHTML, "line-height: 60px;")
	assert.Contains(t, resultHTML, "padding-top: 0;")
	assert.Contains(t, resultHTML, "padding-bottom: 5px")
}

func TestSpecificity(t *testing.T) {
	html := `<html>
        <head>
        <title>Title</title>
        <style type="text/css">
		table.bar-chart td.bar-area {
			padding: 10px;
		}
		table { width: 91%; }
		table { width: 92%; }
		table { width: 93%; }
		table { width: 94%; }
		table { width: 95%; }
		table { width: 96%; }
		table { width: 97%; }
		table.bar-chart td {
			padding: 5px;
		}
        </style>
        </head>
        <body>
		<table class="bar-chart">
			<tr><td>1</td></tr>
			<tr><td class="bar-area">2</td></tr>
		</table>
        </body>
        </html>`

	p, err := NewPremailerFromString(html, NewOptions())
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, `<tr><td style="padding:5px">1</td></tr>`)
	assert.Contains(t, resultHTML, `<tr><td class="bar-area" style="padding:10px">2</td></tr>`)
}

func TestRetainsMsoConditionalComment(t *testing.T) {
	html := `<html>
	<head>
	</head>
	<body>
	<!--[if mso]><style>.body {font-size: 16px;}</style><![endif]-->
	</body>
	</html>`

	p, err := NewPremailerFromString(html, NewOptions())
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, `<!--[if mso]><style>.body {font-size: 16px;}</style><![endif]-->`)
}

func TestRetainsComments(t *testing.T) {
	html := `<html>
	<head>
	</head>
	<body>
	<!-- Comment containing brackets < > -->
	</body>
	</html>`

	p, err := NewPremailerFromString(html, NewOptions())
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, `<!-- Comment containing brackets < > -->`)
}

func TestRetainsLinkText(t *testing.T) {
	htmlString := `<!DOCTYPE html><html><head>
		<style>
		p { color: red; font-size: 14px; font-family: 'Arial'; }
		.header { background-color: blue; }
		</style>
	</head>
	<body>
		<p>This is a test paragraph & more</p>
		<a href="https://example.com?a=1&b=2">
			Click https://example.com?a=1&b=2
		</a>
		<div class="header">
			<div>Header Content & more</div>
		</div>
	</body></html>`

	options := NewOptions()
	options.UnescapedTextNode = true

	p, err := NewPremailerFromString(htmlString, options)
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, `Click https://example.com?a=1&b=2`)
	assert.Contains(t, resultHTML, `https://example.com?a=1&amp;b=2`)
	assert.Contains(t, resultHTML, `This is a test paragraph & more`)
	assert.Contains(t, resultHTML, `Header Content & more`)
}
