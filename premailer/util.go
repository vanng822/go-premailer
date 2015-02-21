package premailer

import (
	"github.com/vanng822/gocssom/cssom"
)

func copyRule(selector string, rule *cssom.CSSRule) *cssom.CSSRule {
	// copy rule for each selector
	copiedStyle := cssom.CSSStyleRule{SelectorText: selector, Styles: rule.Style.Styles}
	copiedRule := &cssom.CSSRule{Type: rule.Type, Style: copiedStyle}
	return copiedRule
}