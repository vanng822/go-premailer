package premailer

import (
	"github.com/vanng822/css"
)

func copyRule(selector string, rule *css.CSSRule) *css.CSSRule {
	// copy rule for each selector
	styles := make([]*css.CSSStyleDeclaration, 0)
	for _, s := range rule.Style.Styles {
		styles = append(styles, css.NewCSSStyleDeclaration(s.Property, s.Value.Text(), s.Important))
	}
	copiedStyle := css.CSSStyleRule{Selector: css.NewCSSValueString(selector), Styles: styles}
	copiedRule := &css.CSSRule{Type: rule.Type, Style: copiedStyle}
	return copiedRule
}

func makeRuleImportant(rule *css.CSSRule) string {
	// this for using Text() which has nice sorted props
	for _, s := range rule.Style.Styles {
		s.Important = true
	}
	return rule.Style.Text()
}
