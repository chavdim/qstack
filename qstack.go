package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"os"
)
//string slicing 
func getFromTo(text string, from string,to string) string {
	i1 := strings.Index(text, from)
	i2 := strings.Index(text, to)
	sub := text[i1:i2]
	return sub
}
func getFromToStartingFrom(text string, from string,to string,) string {
	i1 := strings.Index(text, from)
	i2 := strings.Index(text[i1:], to) + i1
	sub := text[i1:i2]
	return sub
}
func prepareAnswer(ans string) string {
	r := strings.Replace(ans, "<code>", "\nCODE---------------------------------------\n", -1)
	r = strings.Replace(r, "</code>", "\nENDCODE------------------------------------\n", -1)
	r = strings.Replace(r, "<p>", "", -1)
	r = strings.Replace(r, "</p>", "", -1)
	r = strings.Replace(r, "<pre>", "", -1)
	r = strings.Replace(r, "</pre>", "", -1)
	return r
}
func main() {
	//Get args
    args := os.Args[1:]
    searchQuery := strings.Replace(strings.Join(args[:]," "), " ", "+", -1)
  	//
    fmt.Println("Searching for: "+searchQuery)

    //Find top stackoverflow resutl
	url :="https://search.yahoo.com/search?p="+searchQuery
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	html := string(byteArray[:])
	stackUrl := getFromToStartingFrom(html,"http://stackoverflow.com","\"")
	//
	fmt.Println("Source: "+stackUrl)
	//Get stackoverflow page from to result url
	url = stackUrl
	resp, err = http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	byteArray, err = ioutil.ReadAll(resp.Body)
	html = string(byteArray[:])
	//Slice and display results
	fmt.Println("ANSWER##################################################")
	fmt.Println("########################################################")
	sub := getFromToStartingFrom(html,"class=\"answercell\"","class=\"fw\"")
	answer := getFromTo(sub, "<div" ,  "</div>")
	answer = prepareAnswer(answer[39:])
	fmt.Println(answer)

	fmt.Println("QUESTION################################################")
	fmt.Println("########################################################")
	sub = getFromTo(html,"class=\"question-hyperlink\"","id=\"mainbar\"")
	question := getFromTo(sub, "class=\"question-hyperlink\">" ,  "</a>")
	fmt.Println(question[27:])	
}
