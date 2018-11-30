package main

import (
	"fmt"
	"regexp"
	"strings"
)

var domElements = [...]string{
	"a", "abbr", "address", "area", "article", "aside", "audio", "b", "base", "bdo", "blockquote", "body", "br", "button", "canvas", "caption", "site", "code", "col", "colgroup", "datalist", "dd", "del", "details", "dfn", "dialog", "div", "dl", "dt", "em", "embed", "fieldset", "figcaption", "figure", "footer", "form", "head", "header", "h1", "h2", "h3", "h4", "h5", "h6", "hr", "html", "i", "iframe", "img", "ins", "input", "kbd", "label", "legend", "li", "link", "map", "mark", "menu", "menuitem", "meta", "meter", "nav", "object", "ol", "optgroup", "p", "param", "pre", "progress", "q", "s", "samp", "script", "section", "select", "small", "source", "span", "strong", "style", "sub", "summary", "sup", "table", "td", "th", "tr", "tetarea", "time", "title", "track", "u", "ul", "var", "video",
}

var symbols = map[string]string{
	".": "class",
	"#": "id",
}

/* steps
1. spit to array by spaces
2. check special word if there is check if has # or . or none | on false search only by class or id
3. search in text if its found, go to next array item
4. rinse & repeat
*/

// SearchParser does something amazing
func SearchParser(search string) string {

	// split to array
	separated := strings.Split(search, " ")
	// steps := len(separated)

	for _, e := range separated {
		if checkSpecial(e) {
			// memory problem
			special, restSpecial := getSpecial(e)
			if checkIdentifier(restSpecial) {
				identifier, restIdentifier := getIdentifier(restSpecial)
				fmt.Println(special)
				fmt.Println(restSpecial)
				fmt.Println(identifier)
				fmt.Println(restIdentifier)
			} else {
				fmt.Println(e)
			}
		} else if checkIdentifier(e) {
			identifier, restIdentifier := getIdentifier(e)
			fmt.Println(identifier)
			fmt.Println(restIdentifier)
		} else {
			fmt.Println(e)
		}
	}

	// checkIdentifierPattern := fmt.Sprintf("^[\\#\\.](.*)")

	// startPattern := "^[\\.\\#](.*)"
	// if regexp.MatchString(startPattern, search) {

	// } else {

	// }
	// patternClass := "\\.?(.*)"
	// patternID := "\\#?(.*)"
	// check := regexp.MatchString(pattern, search)
	// seperation, _ := regexp.Compile(pattern)
	// f := r.FindStringSubmatch(s)
	return ""
}

// dirty fix | [96]string
func includes(arr [96]string, s string) bool {
	for _, i := range arr {
		if i == s {
			return true
		}
	}
	return false
}

// regex flags "(?i)" case insensetive
func checkSpecial(text string) bool {
	pat := "^(.*)[\\.\\#]?(.*)"
	comp, _ := regexp.Compile(pat)
	r := comp.FindStringSubmatch(text)
	inc := includes(domElements, r[1])
	return inc
}

func getSpecial(text string) (special string, rest string) {
	identifier := checkIdentifier(text)
	if identifier {
		return text, ""
	} else {
		pat := "^(.*)([\\#\\.].*)?"
		comp, _ := regexp.Compile(pat)
		r := comp.FindStringSubmatch(text)
		return r[1], r[2]
	}
}

func checkIdentifier(text string) bool {
	p := "^([\\#\\.])(.*)"
	result, _ := regexp.MatchString(p, text)
	if result {
		return true
	}
	return false
}

func getIdentifier(text string) (identifier string, rest string) {
	pat := "^([\\#\\.])(.*)"
	comp, _ := regexp.Compile(pat)
	r := comp.FindStringSubmatch(text)
	return r[1], r[2]
}

// func domFinder(text string, searched string) string {

// }

// func searchAttribute(text string, attri string) string {

// }
