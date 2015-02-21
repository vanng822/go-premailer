package premailer

import (
	"github.com/PuerkitoBio/goquery"
)

type elementRules struct {
	element *goquery.Selection
	rules   []*styleRule
}

func (er *elementRules) inline() {
	inline, _ := er.element.Attr("style")
	style := mergeRules(inline, er.rules)
	if style != "" {
		er.element.SetAttr("style", style)
	}
}
