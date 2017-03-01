package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"strings"
	"io/ioutil"
)

func TraverseElementNode(stack []string, n *html.Node) {

	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Printf("%v\n", stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {

		TraverseElementNode(stack, c)
	}

}

// todo : to be fixed, can not print all the texts
func TraverseTextNode(mp map[string]string, n *html.Node) {
	if n.Type == html.TextNode {

		text := strings.TrimSpace(n.Data)
		if text != "" {
			// generate ParentElement: Text pair
			mp[n.Parent.Data] = text
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			// don't descend into script or style element
			// since their contents are not visible in a web browser
			if c.Data == "script" || c.Data == "style" {
				continue
				TraverseTextNode(mp, c)
			}
		}

	}
}

func TraverseTextNode2(rd io.Reader) {

	//s := strings.NewReader("<body><h2>Running</h2><h1>Training</h1></body>")
	// tz := html.NewTokenizer(s)
	tz := html.NewTokenizer(rd)

	for {
		tk := tz.Next()
		if tk == html.ErrorToken{
			return
		}
		if tk == html.TextToken {
			text := strings.TrimSpace(string(tz.Text()))
			if text != ""{
				fmt.Println(text)
			}

		}

	}

}

func TraverseANode(links []string, n *html.Node) []string {
	// get `a` elementnode type
	if n.Type == html.ElementNode && n.Data == "a" {

		for _, j := range n.Attr {
			if j.Key == "href" {
				links = append(links, j.Val)

			}
		}

	}

	// after dealing with current node, traverse its children
	for ch := n.FirstChild; ch != nil; ch = ch.NextSibling {
		links = TraverseANode(links, ch)
	}

	return links

}


func GetBodyAndNode(requestUrl string) (io.ReadCloser, *html.Node){

	resp, err := http.Get(requestUrl)
	if err != nil {
		log.Fatalf("%v while requesting %v", err, requestUrl)
	}

	node, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("%v while parsing response from %v", err, requestUrl)
	}
	return  resp.Body,node

}


func TraverseElementNode_test(url string){

	body, node := GetBodyAndNode(url)
	TraverseElementNode(nil,node)
	body.Close()
}


func TraverseANode_test(url string){
	body, node := GetBodyAndNode(url)
	gatheredLinks := TraverseANode(nil, node)
	for _, i := range gatheredLinks {
		fmt.Printf("%v\n", url+i)
	}
	body.Close()

}



func TraverseTextNode_test(url string){
	body, node := GetBodyAndNode(url)
	WhoContainsText:= make(map[string]string)
	TraverseTextNode(WhoContainsText, node)
	for x, y := range WhoContainsText {
		fmt.Printf("%v:\t\t\t%v\n",x,y)
	}
	body.Close()
}

func TraverseTextNode2_test(url string){
	// the reason I don't use GetBodyAndNode
	// is that the body returned by GetBodyAndNode is
	// already used by html.parse,its not the response
	// body that I need. I have to regenerate it! see
	// function RespBody_test for details
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("%v while requesting %v", err, url)
	}
	TraverseTextNode2(resp.Body)
	resp.Body.Close()

}


func RespBody_test(url string){

	body, _ := GetBodyAndNode(url)
	contents,err := ioutil.ReadAll(body)
	fmt.Println("\033[33m",strings.Repeat("*",180),"\033[0m")
	fmt.Println("\033[32m",string(contents),"\033[0m")
	body.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("%v while requesting %v", err, url)
	}
        contents, err = ioutil.ReadAll(resp.Body)
	fmt.Println("\033[33m",strings.Repeat("*",180),"\033[0m")
	fmt.Println("\033[31m",string(contents),"\033[0m")
	fmt.Println("\033[33m",strings.Repeat("*",180),"\033[0m")
	resp.Body.Close()
}

func main() {
	p := fmt.Println
	//var funcs []func(string)
	var funcs = []func(string){TraverseANode_test,
		TraverseElementNode_test,
		TraverseTextNode2_test,
		RespBody_test,
	        TraverseTextNode_test}
	for _, f := range funcs{
		p(strings.Repeat("*",180))
		f("http://www.imdb.com")
	}
	//fmt.Println(strings.Repeat("*",180))
	//TraverseANode_test("http://www.imdb.com")
	//
	//fmt.Println(strings.Repeat("*",180))
	//TraverseElementNode_test("http://www.imdb.com")
	//
	//fmt.Println(strings.Repeat("*",180))
	//TraverseTextNode_test("http://www.imdb.com")
	//
	//fmt.Println(strings.Repeat("*",180))
	//TraverseTextNode2_test("http://www.imdb.com")
	//
	//fmt.Println(strings.Repeat("*",180))
	//RespBody_test("http://www.imdb.com")
}
