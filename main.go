package main

//  html scrapper
// turn into a api maybe

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Icypop struct {
	URI  string
	Body string
}

func (s *Icypop) getBody() {
	resp, err := http.Get(s.URI)
	if err != nil {
		log.Fatalln(err)
	}
	bodyBytes, errByte := ioutil.ReadAll(resp.Body)
	if errByte != nil {
		log.Fatalln(errByte)
	}
	s.Body = string(bodyBytes)

}

func main() {
	first := Icypop{URI: "http://tureng.com/en/turkish-english"}
	first.getBody()

}
