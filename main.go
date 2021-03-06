package main

//  html scrapper
// turn into a api maybe

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// Icypop HTML Parser
type Icypop struct {
	URI  string
	Body string
}

func (i *Icypop) getBody() {
	resp, err := http.Get(i.URI)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, errByte := ioutil.ReadAll(resp.Body)
	if errByte != nil {
		log.Fatalln(errByte)
	}
	i.Body = string(bodyBytes)
}

func (i *Icypop) search(searched string) []string {
	// seperate search for h1..h6 | a | div etc...
	r, err := regexp.Compile("(?i)^a[\\.\\#]?.*?$")
	if err != nil {
		log.Fatalln(err)
	}
	found := r.MatchString(searched)
	if found {
		return searchLink(i.Body)
	}
	return nil
}

func searchLink(s string) []string {
	r, _ := regexp.Compile("(?i)<a.*>(.*)</a>")
	found := r.FindAllString(s, -1)
	return found
}

func main() {
	// first := Icypop{URI: "http://tureng.com/en/turkish-english"}
	// first.getBody()
	// list := first.search("a.a")
	// fmt.Println(list)

	data, _ := ioutil.ReadFile("site.html")
	s := string(data)

	search := "div span.item-change span.arrow-down"
	// search := "div.doviz-column3 div.column3-row2"

	t1 := Parser{s}
	r := t1.SearchElement(search)
	// t := t1.GetText(r)
	for _, i := range r {
		fmt.Println(i)
	}
}

/*
--zhe Plan--

html -> search term -> get href/ get src / get text

*/
