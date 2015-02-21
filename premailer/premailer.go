package premailer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/vanng822/gocssom/cssom"
	"strconv"
	"strings"
)

type Premailer struct {
	doc       *goquery.Document
	elIdAttr  string
	elements  map[int]*elementRules
	rules     []*cssom.CSSRule
	allRules  []*cssom.CSSRule
	elementId int
}

func NewPremailer(doc *goquery.Document) *Premailer {
	premailer := Premailer{}
	premailer.doc = doc
	premailer.rules = make([]*cssom.CSSRule, 0)
	premailer.allRules = make([]*cssom.CSSRule, 0)
	premailer.elements = make(map[int]*elementRules)
	premailer.elIdAttr = "premailer-el-id"
	return &premailer
}

func (premailer *Premailer) sortRules() {
	normalRules := make([]*cssom.CSSRule, 0)
	importantRules := make([]*cssom.CSSRule, 0)

	for _, rule := range premailer.allRules {
		for _, s := range rule.Style.Styles {
			if s.Important == 1 {
				importantRules = append(importantRules, rule)
			} else {
				normalRules = append(normalRules, rule)
			}
		}
	}
	premailer.rules = append(premailer.rules, normalRules...)
	premailer.rules = append(premailer.rules, importantRules...)
}

func (premailer *Premailer) collectRules() {
	premailer.doc.Find("style").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Text())
		//fmt.Println(s.Nodes)
		ss := cssom.Parse(s.Text())
		r := ss.GetCSSRuleList()
		premailer.allRules = append(premailer.allRules, r...)
	})
}

func (premailer *Premailer) collectElements() {
	for _, rule := range premailer.rules {
		fmt.Println(rule.Type, rule.Style.SelectorText)
		if rule.Type == cssom.MEDIA_RULE {
			continue
		}

		selectors := strings.Split(rule.Style.SelectorText, ",")
		for _, selector := range selectors {
			if strings.Contains(selector, ":") {
				continue
			}
			fmt.Println(selector)
			premailer.doc.Find(selector).Each(func(i int, s *goquery.Selection) {
				if val, exist := s.Attr(premailer.elIdAttr); exist {
					fmt.Println("HIT", val)
					id, _ := strconv.Atoi(val)
					premailer.elements[id].rules = append(premailer.elements[id].rules, rule)
				} else {
					s.SetAttr(premailer.elIdAttr, strconv.Itoa(premailer.elementId))
					rules := make([]*cssom.CSSRule, 0)
					rules = append(rules, rule)
					premailer.elements[premailer.elementId] = &elementRules{element: s, rules: rules}
					premailer.elementId += 1
				}
			})
		}
	}
}

func (premailer *Premailer) applyInline() {
	for _, element := range premailer.elements {
		element.inline()
		element.element.RemoveAttr(premailer.elIdAttr)
	}
}

func (premailer *Premailer) Transform() (string, error) {
	premailer.collectRules()
	premailer.sortRules()
	premailer.collectElements()
	premailer.applyInline()
	return premailer.doc.Html()
}
