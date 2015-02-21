package premailer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type elementRules struct {
	element *goquery.Selection
	rules   []*styleRule
}

func (er *elementRules) inline() {
	//inline, _ := er.element.Attr("style")
	styles := make(map[string]string)
	for _, rule := range er.rules {
		for prop, s := range rule.styles {
			styles[prop] = s.Value
		}
	}
	final := make([]string, 0)
	for p, v := range styles {
		final = append(final, fmt.Sprintf("%s:%s", p, v))
		er.style_to_basic_html_attribute(p, v)
	}
	style := strings.Join(final, ";")
	if style != "" {
		er.element.SetAttr("style", style)
	}

}

func (er *elementRules) style_to_basic_html_attribute(prop, value string) {
	switch prop {
	case "text-align":
		er.element.SetAttr("align", value)
	case "vertical-align":
		er.element.SetAttr("valign", value)
	case "background-color":
		er.element.SetAttr("bgcolor", value)
	case "width":
		fallthrough
	case "height":
		if strings.HasSuffix(value, "px") {
			value = value[:len(value)-2]
		}
		er.element.SetAttr(prop, value)
	}
}
