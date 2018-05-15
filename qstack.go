package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"os"
	"errors"
	"github.com/zserge/webview"
)
//string slicing 
func getFromTo(text string, from string,to string) string {
	i1 := strings.Index(text, from)
	i2 := strings.Index(text, to)
	sub := text[i1:i2]
	return sub
}
func getFromToStartingFrom(text string, from string,to string,) (string, error) {
	i1 := strings.Index(text, from)
	if i1==-1 {
		err := errors.New("no matches found, exiting.")
		return "", err 
	}
	i2 := strings.Index(text[i1:], to) + i1	
	sub := text[i1:i2]
	return sub, nil
}
//
func main() {
	//Get args
    args := os.Args[1:]
    searchQuery := "stackoverflow+"
    searchQuery += strings.Replace(strings.Join(args[:]," "), " ", "+", -1)
    if len(args) == 0 {
    	fmt.Println("Please include search term. Example: qstack css add font")
    	return
    }
  	//
    fmt.Println("Searching for: "+searchQuery)
    //Find top stackoverflow result
	url :="https://search.yahoo.com/search?p="+searchQuery
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	html := string(byteArray[:])
	stackUrl, err := getFromToStartingFrom(html,"https://stackoverflow.com/questions","\"")
	if err!=nil{
		log.Fatal(err)
	}
	// displays result in webview
	webview.Open("stackoverflow",
		stackUrl, 800, 600, true)
}
