package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/rocketlaunchr/google-search"
)
func getFromToStartingFrom(text string, from string,to string,) (string, error) {
	i1 := strings.Index(text, from)
	if i1==-1 {
		err := errors.New("No matches found, exiting.")
		return "", err
	}
	i2 := strings.Index(text[i1:], to) + i1
	sub := text[i1:i2]
	return sub, nil
}
//modify answer section
func prepareAnswer(ans string) string {
	r := strings.Replace(ans, "<code>", "\nCODEBLOCK---------------------------------------\n", -1)
	r = strings.Replace(r, "</code>", "\n------------------------------------------------\n", -1)
	r = strings.Replace(r, "<p>", "", -1)
	r = strings.Replace(r, "</p>", "\n", -1)
	r = strings.Replace(r, "<pre>", "", -1)
	r = strings.Replace(r, "</pre>", "\n", -1)
	return r
}

//
func main() {
	//Get args
	args := os.Args[1:]
	searchQuery := "stackoverflow "+strings.Join(args[:], " ")
	if len(args) == 0 {
		fmt.Println("Please include search term. Example: qstack css add font")
		return
	}

	fmt.Println("Searching for: " + searchQuery)

	ctx := context.Background()
	result, err := googlesearch.Search(ctx, searchQuery)
	var firstStackUrlFound = ""
	for _, s := range result {
		if strings.Contains(s.URL, "stackoverflow") {
			firstStackUrlFound = s.URL
			break
		}
	}
	if firstStackUrlFound == "" {
		log.Fatal("No stack overflow page found")
	}

	fmt.Println("Source: " + firstStackUrlFound)
	//Get stackoverflow page from to result url
	resp, err := http.Get(firstStackUrlFound)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var html = string(byteArray[:])
	//Slice and display results
	fmt.Println("Question ##################################################")
	var lastSlash = strings.LastIndex(firstStackUrlFound, "/")
	fmt.Println(strings.Replace(firstStackUrlFound[lastSlash+1:], "-"," ", -1))

	fmt.Println("Answer ##################################################")
	//fmt.Println(html)

	answer, err := getFromToStartingFrom(html, "class=\"answercell", "class=\"mt24\"")
	if err != nil {
		log.Fatal(err)
	}

	// Custom slicing according to response dom
	var i1 = strings.Index(answer, "itemprop=\"text\">")

	answer = prepareAnswer(answer[i1+17 : (len(answer) - 18)])
	fmt.Println(answer)
}