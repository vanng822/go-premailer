package premailer

import (
	"strings"
	"regexp"
	)

// https://developer.mozilla.org/en-US/docs/Web/CSS/Specificity
// https://developer.mozilla.org/en-US/docs/Web/CSS/Reference#Selectors

type specificity struct {
	important   int
	id_count    int
	class_count int
	type_count  int
	attr_count  int
}


var _type_selector_regex = regexp.MustCompile("(^|\\s)\\w")

func makeSpecificity(important int, selector string) *specificity {
	spec := specificity{}
	// determine values for priority
	spec.important = important
	spec.id_count = strings.Count(selector, "#")
	spec.class_count = strings.Count(selector, ".")
	spec.type_count = len(_type_selector_regex.FindAllString(selector, -1))
	return &spec
}

type bySpecificity []*styleRule

func (bs bySpecificity) Len() int {
	return len(bs)
}
func (bs bySpecificity) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}
func (bs bySpecificity) Less(i, j int) bool {
	if bs[i].specificity.important < bs[j].specificity.important {
		return true
	}
	
	if bs[i].specificity.id_count < bs[j].specificity.id_count {
		return true
	}
	if bs[i].specificity.class_count < bs[j].specificity.class_count {
		return true
	}
	
	if bs[i].specificity.type_count < bs[j].specificity.type_count {
		return true
	}
	
	return false
}
