package premailer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/vanng822/gocssom/cssom"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Premailer interface {
	Transform() (string, error)
}

type styleRule struct {
	specificity *specificity
	selector    string
	styles      map[string]*cssom.CSSStyleDeclaration
}

type premailer struct {
	doc       *goquery.Document
	elIdAttr  string
	elements  map[int]*elementRules
	rules     []*styleRule
	leftover  []*cssom.CSSRule
	allRules  [][]*cssom.CSSRule
	elementId int
	processed bool
}

func NewPremailer(doc *goquery.Document) Premailer {
	pr := premailer{}
	pr.doc = doc
	pr.rules = make([]*styleRule, 0)
	pr.allRules = make([][]*cssom.CSSRule, 0)
	pr.leftover = make([]*cssom.CSSRule, 0)
	pr.elements = make(map[int]*elementRules)
	pr.elIdAttr = "pr-el-id"
	return &pr
}

func copyRule(selector string, rule *cssom.CSSRule) *cssom.CSSRule {
	// copy rule for each selector
	copiedStyle := cssom.CSSStyleRule{SelectorText: selector, Styles: rule.Style.Styles}
	copiedRule := &cssom.CSSRule{Type: rule.Type, Style: copiedStyle}
	return copiedRule
}

func (pr *premailer) sortRules() {
	ruleIndexCounter := 1
	for ruleSetIndex, rules := range pr.allRules {
		if rules == nil {
			continue
		}

		for _, rule := range rules {
			if rule.Type == cssom.MEDIA_RULE {
				pr.leftover = append(pr.leftover, rule)
				continue
			}
			normalStyles := make(map[string]*cssom.CSSStyleDeclaration)
			importantStyles := make(map[string]*cssom.CSSStyleDeclaration)

			for prop, s := range rule.Style.Styles {
				fmt.Println(s.Value)
				if s.Important == 1 {
					importantStyles[prop] = s
				} else {
					normalStyles[prop] = s
				}
			}

			selectors := strings.Split(rule.Style.SelectorText, ",")
			for _, selector := range selectors {

				if strings.Contains(selector, ":") {
					// cause longer css
					pr.leftover = append(pr.leftover, copyRule(selector, rule))
					continue
				}
				if strings.Contains(selector, "*") {
					// keep this?
					pr.leftover = append(pr.leftover, copyRule(selector, rule))
					continue
				}
				// TODO: Calculate specificity https://developer.mozilla.org/en-US/docs/Web/CSS/Specificity
				// instead if this and sort on it
				if len(normalStyles) > 0 {
					pr.rules = append(pr.rules, &styleRule{makeSpecificity(0, ruleSetIndex, ruleIndexCounter, selector), selector, normalStyles})
					ruleIndexCounter += 1
				}
				if len(importantStyles) > 0 {
					pr.rules = append(pr.rules, &styleRule{makeSpecificity(1, ruleSetIndex, ruleIndexCounter, selector), selector, importantStyles})
					ruleIndexCounter += 1
				}
			}
		}
	}
	// TODO sort by specificity
	//pr.rules = append(pr.rules, normalStyles...)
	//pr.rules = append(pr.rules, importantStyles...)
	sort.Sort(bySpecificity(pr.rules))
}

func (pr *premailer) collectRules() {
	var wg sync.WaitGroup
	pr.doc.Find("style").Each(func(i int, s *goquery.Selection) {
		wg.Add(1)
		pr.allRules = append(pr.allRules, nil)
		go func() {
			defer wg.Done()
			ss := cssom.Parse(s.Text())
			r := ss.GetCSSRuleList()
			pr.allRules[i] = r
		}()
	})
	wg.Wait()

}

func (pr *premailer) collectElements() {
	for _, rule := range pr.rules {
		fmt.Println(rule.selector, rule.specificity)

		pr.doc.Find(rule.selector).Each(func(i int, s *goquery.Selection) {
			if val, exist := s.Attr(pr.elIdAttr); exist {
				fmt.Println("HIT", val)
				id, _ := strconv.Atoi(val)
				pr.elements[id].rules = append(pr.elements[id].rules, rule)
			} else {
				s.SetAttr(pr.elIdAttr, strconv.Itoa(pr.elementId))
				rules := make([]*styleRule, 0)
				rules = append(rules, rule)
				pr.elements[pr.elementId] = &elementRules{element: s, rules: rules}
				pr.elementId += 1
			}
		})

	}
}

func (pr *premailer) applyInline() {
	for _, element := range pr.elements {
		element.inline()
		element.element.RemoveAttr(pr.elIdAttr)
	}
}

func (pr *premailer) Transform() (string, error) {
	if !pr.processed {
		pr.collectRules()
		pr.sortRules()
		pr.collectElements()
		pr.applyInline()
		pr.processed = true
	}
	return pr.doc.Html()
}
