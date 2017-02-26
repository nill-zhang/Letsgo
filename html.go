package main

import ("fmt"
	"golang.org/x/net/html"
	"net/http"
	"log"
)

func ExtractLinks(requestUrl string){

	rep, err := http.Get(requestUrl)
	if err != nil{
		// equivalent of fmt.Printf(fmt,...);os.Exit(1)
		log.Fatalf("%v while connecting to %v", err, requestUrl)

	}
	node, err := html.Parse(rep.Body)
	// close resources
	defer rep.Body.Close()
	if err != nil {
		log.Fatalf("%v while parsing response from %v", err, requestUrl)
	}
        // Note that I use nil to initialize slice parameters
	gatheredLinks := TraverseANode(nil,node)
	TraverseElementNode(nil, node)
	for _,j :=range gatheredLinks{
		fmt.Printf("%v\n",requestUrl+j)
	}

}

func TraverseElementNode(stack []string, n *html.Node){

	if n.Type == html.ElementNode{
		stack = append(stack, n.Data)
		fmt.Printf("%v\n",stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling{

		TraverseElementNode(stack, c)
	}



}

func TraverseANode(links []string, n *html.Node)[]string{
	// get `a` elementnode type
	if n.Type == html.ElementNode && n.Data == "a"{

		for _,j := range n.Attr{
			if j.Key == "href"{
				links = append(links, j.Val)

			}
		}

	}


	// after dealing with current node, traverse its children
	for ch := n.FirstChild; ch != nil;ch = ch.NextSibling{
		links = TraverseANode(links, ch)
	}

	return  links


}


func main(){

	ExtractLinks("http://www.imdb.com")


}
