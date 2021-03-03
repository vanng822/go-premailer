package premailer

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/vanng822/css"
)

type void struct{}

type elementRules struct {
	element           *goquery.Selection
	rules             []*styleRule
	cssToAttributes   bool
	keepBangImportant bool
}

func (er *elementRules) inline() {
	inline, _ := er.element.Attr("style")

	inlineStyles := make([]*css.CSSStyleDeclaration, 0)
	if inline != "" {
		inlineStyles = css.ParseBlock(inline)
	}
	// we collect all occurrences
	orders := make([]string, 0)
	styles := make(map[string]string)
	for _, rule := range er.rules {
		for _, s := range rule.styles {
			prop := s.Property
			styles[prop] = s.Value.Text()
			if er.keepBangImportant && s.Important {
				styles[prop] += " !important"
			}
			orders = append(orders, prop)
		}
	}

	if len(inlineStyles) > 0 {
		for _, s := range inlineStyles {
			prop := s.Property
			styles[prop] = s.Value.Text()
			orders = append(orders, prop)
		}
	}

	// now collect the order of property
	// each prop will be unique and the last one
	// should have the highest priority
	// We could inline them all but this will make output cleaner
	props := make([]string, 0)
	seen := make(map[string]void)
	for i := len(orders) - 1; i >= 0; i-- {
		prop := orders[i]
		if _, ok := seen[prop]; !ok {
			seen[prop] = void{}
			props = append(props, prop)
		}
	}

	final := make([]string, 0, len(styles))
	for i := len(props) - 1; i >= 0; i-- {
		p := props[i]
		v := styles[p]
		final = append(final, fmt.Sprintf("%s:%s", p, v))
		if er.cssToAttributes {
			er.styleToBasicHtmlAttribute(p, v)
		}
	}

	style := strings.Join(final, ";")
	if style != "" {
		er.element.SetAttr("style", style)
	}

}

func (er *elementRules) styleToBasicHtmlAttribute(prop, value string) {
	switch prop {
	case "width":
		fallthrough
	case "height":
		// If we are keeping "!important" in our styles, we need to strip it out
		// here so that the proper px value can be parsed
		if er.keepBangImportant {
			value = strings.Replace(value, " !important", "", 1)
		}
		if strings.HasSuffix(value, "px") {
			value = value[:len(value)-2]
			er.element.SetAttr(prop, value)
		} else if value == "0" {
			er.element.SetAttr(prop, value)
		}
	}
}
