package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	//"bufio"
	"os"
)

func getFromTo(text string, from string,to string) string {
	i1 := strings.Index(text, from)
	i2 := strings.Index(text, to)
	//fmt.Println(i1)
	//fmt.Println(i2)
	sub := text[i1:i2]
	return sub
}
func getFromToStartingFrom(text string, from string,to string,) string {
	i1 := strings.Index(text, from)
	i2 := strings.Index(text[i1:], to) + i1
	//fmt.Println(i1)
	//fmt.Println(i2)
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
/*	
func prepareAnswer(ans string) string {
	toAdd := "CODE-------------------------------------------\n"
	lenOfAddition := len(toAdd)
	fmt.Println(lenOfAddition)
	//i1 := strings.Index(ans, "<code>")
	i1 := strings.Index(ans, "<code>")
	startFrom := 0
	t:=3
	//for i1!=-1{
	for t>0{
		i2 := strings.Index(ans[i1:], "</code>") + i1 
		ans = ans[0:i1+6] + toAdd + ans[i1+6:i2+7] + toAdd + ans[i2+7:]
		 fmt.Println(ans)
		 startFrom=i2+(lenOfAddition*2)+startFrom
		i1 = strings.Index(ans[startFrom:], "<code>") + startFrom
		t-=1
	}
	//i1 := strings.Index(ans, "<code>")
	//i2 := strings.Index(ans[i1:], "</code>") + i1
	//fmt.Println(i1)
	//fmt.Println(i2)
	//r := ans[0:i1+6] + toAdd + ans[i1+6:i2+6] + toAdd + ans[i2+6:]
	return ans
}
*/
func main() {
	
    args := os.Args[1:]
    searchQuery := strings.Replace(strings.Join(args[:]," "), " ", "+", -1)
  
    fmt.Println("Searching for: "+searchQuery)
	/*
	url := "http://stackoverflow.com/questions/38856593/cant-return-from-server-method-with-meteor-settimout"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	html := string(byteArray[:])
	*/
	//url = "https://www.google.co.jp/webhp?sourceid=chrome-instant&ion=1&espv=2&ie=UTF-8#q=golang%20string%20to%20int"
	//url :="https://search.yahoo.com/search?p=golang+length+array"
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


	//url := "http://stackoverflow.com/questions/38856593/cant-return-from-server-method-with-meteor-settimout"
	url = stackUrl
	resp, err = http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	byteArray, err = ioutil.ReadAll(resp.Body)
	html = string(byteArray[:])
	
	/*

	// Open the file.
    f, _ := os.Open("html.txt")
    // Create a new Scanner for the file.
    scanner := bufio.NewScanner(f)
    // Loop over all lines in the file and print them.
    for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
    }
    */
    /*
    pwd, _ := os.Getwd()
    byteArray, err :=  ioutil.ReadFile(pwd+"/work/src/github.com/user/qstack/html.txt")
    if err != nil {
		log.Fatal(err)
	}
	html := string(byteArray[:])
	*/

	//html := " <div class=favicon favicon-sports title=Sports Stack Exchange></div><a href=http:/"


	
	
	sub := getFromTo(html,"class=\"question-hyperlink\"","id=\"mainbar\"")
	question := getFromTo(sub, "class=\"question-hyperlink\">" ,  "</a>")
	

	/*
	i1 = html.index('class="answer accepted-answer"')
	i2 = html.index('class="post-text"',i1)
	i3 = html.index('</div>',i2)
	ans = html[i2:i3]
	*/
	fmt.Println("ANSWER##################################################")
	fmt.Println("########################################################")

	sub = getFromToStartingFrom(html,"class=\"answercell\"","class=\"fw\"")
	answer := getFromTo(sub, "<div" ,  "</div>")
	answer = prepareAnswer(answer[39:])
	fmt.Println(answer)

	fmt.Println("QUESTION################################################")
	fmt.Println("########################################################")
	fmt.Println(question[27:])



	
	
	
}
