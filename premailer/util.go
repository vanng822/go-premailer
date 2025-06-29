package premailer

import (
	"github.com/vanng822/css"
	"golang.org/x/net/html"
)

func copyRule(selector string, rule *css.CSSRule) *css.CSSRule {
	// copy rule for each selector
	styles := make([]*css.CSSStyleDeclaration, 0)
	for _, s := range rule.Style.Styles {
		styles = append(styles, css.NewCSSStyleDeclaration(s.Property, s.Value.Text(), s.Important))
	}
	copiedStyle := css.CSSStyleRule{Selector: css.NewCSSValue(selector), Styles: styles}
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

func makeUnsafeRawTextNode(s *html.Node) {
	for c := s.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			c.Type = html.RawNode
			continue
		}
		makeUnsafeRawTextNode(c)
	}
}
