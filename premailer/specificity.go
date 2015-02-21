package premailer

import (
	"strings"
	"regexp"
	)

// https://developer.mozilla.org/en-US/docs/Web/CSS/Specificity
// https://developer.mozilla.org/en-US/docs/Web/CSS/Reference#Selectors

type specificity struct {
	important   int
	idCount    int
	classCount int
	typeCount  int
	attrCount  int
	ruleSetIndex int
	ruleIndex int
}


var _type_selector_regex = regexp.MustCompile("(^|\\s)\\w")

func makeSpecificity(important, ruleSetIndex, ruleIndex int, selector string) *specificity {
	spec := specificity{}
	// determine values for priority
	spec.important = important
	spec.idCount = strings.Count(selector, "#")
	spec.classCount = strings.Count(selector, ".")
	spec.typeCount = len(_type_selector_regex.FindAllString(selector, -1))
	spec.ruleSetIndex = ruleSetIndex
	spec.ruleIndex = ruleIndex
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
	
	if bs[i].specificity.idCount < bs[j].specificity.idCount {
		return true
	}
	
	if bs[i].specificity.classCount < bs[j].specificity.classCount {
		return true
	}
	
	if bs[i].specificity.attrCount < bs[j].specificity.attrCount {
		return true
	}
	
	if bs[i].specificity.typeCount < bs[j].specificity.typeCount {
		return true
	}
	
	if bs[i].specificity.ruleSetIndex < bs[j].specificity.ruleSetIndex {
		return true
	}
	
	if bs[i].specificity.ruleIndex < bs[j].specificity.ruleIndex {
		return true
	}
	
	return false
}
