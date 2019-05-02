package premailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicHTMLFromFile(t *testing.T) {
	p, err := NewPremailerFromFile("data/markup_test.html", nil)
	assert.Nil(t, err)
	resultHTML, err := p.Transform()
	assert.Nil(t, err)

	assert.Contains(t, resultHTML, "<h1 style=\"width:50px;color:red\" width=\"50\">Hi!</h1>")
	assert.Contains(t, resultHTML, "<h2 style=\"vertical-align:top\">There</h2>")
	assert.Contains(t, resultHTML, "<h3 style=\"text-align:right\">Hello</h3>")
	assert.Contains(t, resultHTML, "<p><strong style=\"text-decoration:none\">Yes!</strong></p>")
	assert.Contains(t, resultHTML, "<div style=\"background-color:green\">Green color</div>")
}

func TestFromFileNotFound(t *testing.T) {
	p, err := NewPremailerFromFile("data/blablabla.html", nil)
	assert.NotNil(t, err)
	assert.Nil(t, p)
}
