package premailer

import (
	"fmt"
	"github.com/vanng822/gocssom/cssom"
	"strings"
)

func mergeRules(inline string, rules []*cssom.CSSRule) string {
	style := make(map[string]string)
	for _, rule := range rules {
		for prop, s := range rule.Style.Styles {
			style[prop] = s.Value
		}
	}
	final := make([]string, 0)
	for p, v := range style {
		final = append(final, fmt.Sprintf("%s:%s", p, v))
	}
	return strings.Join(final, ";")
}
