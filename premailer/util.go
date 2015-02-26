package premailer

import (
	"github.com/vanng822/css"
)

func copyRule(selector string, rule *css.CSSRule) *css.CSSRule {
	// copy rule for each selector
	copiedStyle := css.CSSStyleRule{SelectorText: selector, Styles: rule.Style.Styles}
	copiedRule := &css.CSSRule{Type: rule.Type, Style: copiedStyle}
	return copiedRule
}

func makeRuleImportant(rule *css.CSSRule) string {
	// this for using Text() which has nice sorted props
	for _, s := range rule.Style.Styles {
		s.Important = 1
	}
	return rule.Style.Text()
}