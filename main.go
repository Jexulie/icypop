package main

//  html scrapper
// turn into a api maybe

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

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
	r, err := regexp.Compile("^a[\\.\\#]?.*?$")
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
	r, _ := regexp.Compile("<a.*>(.*)</a>")
	found := r.FindAllString(s, -1)
	return found
}

func main() {
	// first := Icypop{URI: "http://tureng.com/en/turkish-english"}
	// first.getBody()
	// list := first.search("a.a")
	// fmt.Println(list)

	// t1 := "h1.green a p.big"
	t2 := "div.header h2#jumbotron a.lastlink"
	t3 := "br#seperator"
	t4 := "h3#header a.blue"

	// h1 := "<h1 class=\"green\"><a href=\"www.google.com\"><p class=\"big\">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Nulla, ratione!</p></a></h1>"
	h2 := "<div class=\"header\"><h2 id=\"jumbotron\"><a href=\"#\" class=\"lastlink\">site is empty</a></h2></div>"
	h3 := "<br id=\"seperator\">"
	h4 := "<h3 id=\"header\"><a href=\"#\" class=\"blue\">omg its empty</a></h3>"

	// SearchParser(t1, h1)
	SearchParser(t2, h2)
	SearchParser(t3, h3)
	SearchParser(t4, h4)
}
