package premailer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/vanng822/gocssom/cssom"
	"strconv"
	"strings"
)

type Premailer interface {
	Transform() (string, error)
}

type premailer struct {
	doc       *goquery.Document
	elIdAttr  string
	elements  map[int]*elementRules
	rules     []*cssom.CSSRule
	allRules  []*cssom.CSSRule
	elementId int
}

func NewPremailer(doc *goquery.Document) Premailer {
	pr := premailer{}
	pr.doc = doc
	pr.rules = make([]*cssom.CSSRule, 0)
	pr.allRules = make([]*cssom.CSSRule, 0)
	pr.elements = make(map[int]*elementRules)
	pr.elIdAttr = "pr-el-id"
	return &pr
}

func (pr *premailer) sortRules() {
	normalRules := make([]*cssom.CSSRule, 0)
	importantRules := make([]*cssom.CSSRule, 0)

	for _, rule := range pr.allRules {
		for _, s := range rule.Style.Styles {
			if s.Important == 1 {
				importantRules = append(importantRules, rule)
			} else {
				normalRules = append(normalRules, rule)
			}
		}
	}
	pr.rules = append(pr.rules, normalRules...)
	pr.rules = append(pr.rules, importantRules...)
}

func (pr *premailer) collectRules() {
	pr.doc.Find("style").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Text())
		//fmt.Println(s.Nodes)
		ss := cssom.Parse(s.Text())
		r := ss.GetCSSRuleList()
		pr.allRules = append(pr.allRules, r...)
	})
}

func (pr *premailer) collectElements() {
	for _, rule := range pr.rules {
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
			pr.doc.Find(selector).Each(func(i int, s *goquery.Selection) {
				if val, exist := s.Attr(pr.elIdAttr); exist {
					fmt.Println("HIT", val)
					id, _ := strconv.Atoi(val)
					pr.elements[id].rules = append(pr.elements[id].rules, rule)
				} else {
					s.SetAttr(pr.elIdAttr, strconv.Itoa(pr.elementId))
					rules := make([]*cssom.CSSRule, 0)
					rules = append(rules, rule)
					pr.elements[pr.elementId] = &elementRules{element: s, rules: rules}
					pr.elementId += 1
				}
			})
		}
	}
}

func (pr *premailer) applyInline() {
	for _, element := range pr.elements {
		element.inline()
		element.element.RemoveAttr(pr.elIdAttr)
	}
}

func (pr *premailer) Transform() (string, error) {
	pr.collectRules()
	pr.sortRules()
	pr.collectElements()
	pr.applyInline()
	return pr.doc.Html()
}
