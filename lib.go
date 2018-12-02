package main

import (
	"fmt"
	"regexp"
	"strings"
)

// TODO: multi class entry  -> .item-change.color-red

var domElements = [...]string{
	"a", "abbr", "address", "area", "article", "aside", "audio", "b", "base", "bdo", "blockquote", "body", "br", "button", "canvas", "caption", "site", "code", "col", "colgroup", "datalist", "dd", "del", "details", "dfn", "dialog", "div", "dl", "dt", "em", "embed", "fieldset", "figcaption", "figure", "footer", "form", "head", "header", "h1", "h2", "h3", "h4", "h5", "h6", "hr", "html", "i", "iframe", "img", "ins", "input", "kbd", "label", "legend", "li", "link", "map", "mark", "menu", "menuitem", "meta", "meter", "nav", "object", "ol", "optgroup", "p", "param", "pre", "progress", "q", "s", "samp", "script", "section", "select", "small", "source", "span", "strong", "style", "sub", "summary", "sup", "table", "td", "th", "tr", "tetarea", "time", "title", "track", "u", "ul", "var", "video",
}

// not done yet
var elemsNoBrackets = [...]string{
	"input", "img", "br", "hr",
}

var attributes = [...]string{
	"href", "src",
}

var symbols = map[string]string{
	".": "class",
	"#": "id",
}

// regexp patterns
var elementPatt = "(?imU)<%s(.*)>(.*)</%s>"
var singleElemPatt = "(?imU)<%s(.*)/?>"

// fix wildcard wierdness
var searchByIDPatt = "(?im).*%s=\".*%s.*\".*"

// var searchByIDPatt = "(?im).*%s=\"%s\".*"

var checkIDPatt = "(?im)(.*)([\\#\\.])(.*)"

var hrefPatt = "(?imU)href=\"(.*)\""
var srcPatt = "(?imU)src=\"(.*)\""
var textPatt = "(?im)>(.*)<"

// Checks element in an array
func includes(arr []string, s string) bool {
	for _, i := range arr {
		if i == s {
			return true
		}
	}
	return false
}

// Error checker
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

// Parser - parser struct
type Parser struct {
	HTML string
}

// SearchElement search for HTML element
func (p *Parser) SearchElement(search string) []string {
	var result []string
	seperated := strings.Split(search, " ")

	if len(seperated) == 0 {
		panic("Enter search string...")
	} else if len(seperated) > 1 {
		fmt.Println("multi search")
		result = p.multiSearch(seperated)
	} else {
		fmt.Println("single search")
		result = p.singleSearch(seperated[0])
	}
	return result
}

func (p *Parser) multiSearch(words []string) []string {
	text := words
	totalSteps := len(words)
	steps := totalSteps
	var found []string
	for _, word := range text {
		if p.checkForIdentifier(word) {
			stem, idType, id := p.getIdentifier(word)

			// fmt.Println(stem, idType, id)
			// gets same results every loop -- error
			if steps == totalSteps-1 {
				search := p.getElement(stem)
				pat := fmt.Sprintf(searchByIDPatt, idType, id)
				c, err := regexp.Compile(pat)
				checkErr(err)

				var subFound []string

				for _, s := range search {
					if c.MatchString(s) {
						subFound = append(subFound, s)
					}
				}

				steps--
				if steps == 0 {
					return found
				}
				text = subFound
			} else {
				search := p.getElementLoop(text, stem)
				pat := fmt.Sprintf(searchByIDPatt, idType, id)
				c, err := regexp.Compile(pat)
				checkErr(err)

				var subFound []string

				for _, s := range search {
					if c.MatchString(s) {
						subFound = append(subFound, s)
					}
				}

				steps--
				if steps == 0 {
					found = subFound
					return found
				}
				text = subFound
			}
		} else {
			search := p.getElement(word)
			steps--
			if steps == 0 {
				return search
			}
			text = search
		}
	}
	return found
}

func (p *Parser) singleSearch(word string) []string {
	if p.checkForIdentifier(word) {
		stem, idType, id := p.getIdentifier(word)
		search := p.getElement(stem)

		var found []string

		pat := fmt.Sprintf(searchByIDPatt, idType, id)
		c, err := regexp.Compile(pat)
		checkErr(err)

		for _, s := range search {
			if c.MatchString(s) {
				found = append(found, s)
			}
		}
		return found
	}
	// w/o id
	return p.getElement(word)
}

func (p *Parser) getElement(word string) []string {
	var pat string

	if includes(elemsNoBrackets[:], word) {
		pat = fmt.Sprintf(singleElemPatt, word)
	} else {
		pat = fmt.Sprintf(elementPatt, word, word)
	}
	c, err := regexp.Compile(pat)
	checkErr(err)
	res := c.FindAllString(p.HTML, -1)
	return res
}

func (p *Parser) getElementLoop(words []string, stem string) []string {
	var pat string
	text := strings.Join(words, " ")
	if includes(elemsNoBrackets[:], stem) {
		pat = fmt.Sprintf(singleElemPatt, stem)
	} else {
		pat = fmt.Sprintf(elementPatt, stem, stem)
	}
	c, err := regexp.Compile(pat)
	checkErr(err)
	res := c.FindAllString(text, -1)
	return res
}

func (p *Parser) checkForIdentifier(search string) bool {
	c, err := regexp.Compile(checkIDPatt)
	checkErr(err)
	res := c.MatchString(search)
	return res
}

func (p *Parser) getIdentifier(search string) (string, string, string) {
	c, err := regexp.Compile(checkIDPatt)
	checkErr(err)
	res := c.FindStringSubmatch(search)
	return res[1], symbols[res[2]], res[3]
}

func (p *Parser) GetHref(words []string) []string {
	var result []string
	for _, word := range words {
		c, err := regexp.Compile(hrefPatt)
		checkErr(err)
		res := c.FindStringSubmatch(word)
		if len(res) > 0 {
			result = append(result, res[1])
		}
	}
	return result
}

func (p *Parser) GetSrc(words []string) []string {
	var result []string
	for _, word := range words {
		c, err := regexp.Compile(srcPatt)
		checkErr(err)
		res := c.FindStringSubmatch(word)
		if len(res) > 0 {
			result = append(result, res[1])
		}
	}
	return result
}

func (p *Parser) GetText(words []string) []string {
	var result []string
	for _, word := range words {
		c, err := regexp.Compile(textPatt)
		checkErr(err)
		res := c.FindStringSubmatch(word)
		if len(res) > 0 {
			result = append(result, res[1])
		}
	}
	return result
}
