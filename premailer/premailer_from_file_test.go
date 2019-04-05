package premailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicHTMLFromFile(t *testing.T) {
	p, err := NewPremailerFromFile("data/markup_test.html", nil)
	assert.Nil(t, err)
	result_html, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, result_html, "<h1 style=\"color:red;width:50px\" width=\"50\">Hi!</h1>")
	assert.Contains(t, result_html, "<h2 style=\"vertical-align:top\" valign=\"top\">There</h2>")
	assert.Contains(t, result_html, "<h3 style=\"text-align:right\" align=\"right\">Hello</h3>")
	assert.Contains(t, result_html, "<p><strong style=\"text-decoration:none\">Yes!</strong></p>")
	assert.Contains(t, result_html, "<div style=\"background-color:green\" bgcolor=\"green\">Green color</div>")
}

func TestFromFileNotFound(t *testing.T) {
	p, err := NewPremailerFromFile("data/blablabla.html", nil)
	assert.NotNil(t, err)
	assert.Nil(t, p)
}
