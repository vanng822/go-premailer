package premailer

import (
	"fmt"
	"strings"
)

func mergeRules(inline string, rules []*styleRule) string {
	style := make(map[string]string)
	for _, rule := range rules {
		for prop, s := range rule.styles {
			style[prop] = s.Value
		}
	}
	final := make([]string, 0)
	for p, v := range style {
		final = append(final, fmt.Sprintf("%s:%s", p, v))
	}
	return strings.Join(final, ";")
}
