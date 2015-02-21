package premailer

import (
	"github.com/vanng822/gocssom/cssom"
)

type styleRule struct {
	specificity *specificity
	selector    string
	styles      map[string]*cssom.CSSStyleDeclaration
}