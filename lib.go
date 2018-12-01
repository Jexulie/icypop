package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Parser parser struct
type Parser struct {
}

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

/* steps
1. spit to array by spaces
2. check object word if there is check if has # or . or none | on false search only by class or id
3. search in text if its found, go to next array item
4. rinse & repeat
*/

// SearchParser does something amazing
func SearchParser(search string, text string) {

	// split to array
	separated := strings.Split(search, " ")

	// text changes every loop
	// result in the end
	t := text

	for _, e := range separated {
		m := recuvCheck(e)
		rest := domFinder(t, m)
		t = rest
		// fmt.Println(s[2])
	}

	fmt.Println(t)
}

// seperation logic - recursive usage
func recuvCheck(text string) map[string]string {
	element := make(map[string]string)
	if checkAlone(text) {
		// div | a | p goes to parse
		element["object"] = text
		return element
	} else {
		if checkObject(text) {
			object, restObject := getObject(text)
			element["object"] = object

			if checkIdentifier(restObject) {
				identifier, restIdentifier := getIdentifier(restObject)
				element["identifier"] = symbols[identifier]
				element["name"] = restIdentifier
				return element
			} else {
				element["name"] = restObject
				return element
			}
		} else {
			identifier, restIdentifier := getIdentifier(text)
			element["identifier"] = symbols[identifier]
			element["name"] = restIdentifier

			return element
		}
	}
}

// dirty fix | [96]string -- testing array[:]
func includes(arr []string, s string) bool {
	for _, i := range arr {
		if i == s {
			return true
		}
	}
	return false
}

// separates w/o class|id
func checkAlone(text string) bool {
	pat := "^(.*)([\\#\\.].*)"
	result, _ := regexp.MatchString(pat, text)
	return !result
}

// regex flags "(?i)" case insensetive
func checkObject(text string) bool {
	pat := "^(.*)[\\.\\#](.*)"
	comp, _ := regexp.Compile(pat)
	r := comp.FindStringSubmatch(text)
	inc := includes(domElements[:], r[1])
	return inc
}

// gets dom object and rest of the string
func getObject(text string) (object string, rest string) {
	pat := "^(.*)([\\#\\.].*)"
	comp, _ := regexp.Compile(pat)
	r := comp.FindStringSubmatch(text)
	return r[1], r[2]
}

// checks class|id
func checkIdentifier(text string) bool {
	p := "^(.*)([\\#\\.])(.*)"
	result, _ := regexp.MatchString(p, text)
	if result {
		return true
	}
	return false
}

// gets class|id and rest of the string
func getIdentifier(text string) (identifier string, rest string) {
	pat := "^([\\#\\.])(.*)"
	comp, _ := regexp.Compile(pat)
	r := comp.FindStringSubmatch(text)
	return r[1], r[2]
}

// array str for now
func domFinder(text string, searched map[string]string) string {
	// testing <%s>(.*)</%s>
	// 2 <%s(.*)/?>

	// get whole tag instead of text in it
	var htmlBrackets string
	if includes(elemsNoBrackets[:], searched["object"]) {
		htmlBrackets = fmt.Sprintf("<%s(.*)/?>", searched["object"])
	} else {
		htmlBrackets = fmt.Sprintf(">(.*)</%s", searched["object"])
	}

	comp, _ := regexp.Compile(htmlBrackets)
	r := comp.FindStringSubmatch(text)
	if len(r) > 1 {
		return r[1]
	}
	return r[0]
	// 1. object finder
	// 2. identifier finder ?!
	// 3. name finder then get text in it
	// 4. or get attribute text

}

// func searchAttribute(text string, attri string) string {

// }
