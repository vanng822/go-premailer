package premailer

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/vanng822/gocssom/cssom"
)

type ElementRules struct {
	element *goquery.Selection
	rules   []*cssom.CSSRule
}

func (er *ElementRules) Inline() {
	inline, _ := er.element.Attr("style")
	style := MergeRules(inline, er.rules)
	if style != "" {
		er.element.SetAttr("style", style)
	}
}
