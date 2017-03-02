package main
import ("fmt"
	"strings"
	"golang.org/x/net/html"
	"net/http"
	"log"
)
func Expand(s string,f func(string)string)string{

	return strings.Replace(s, "$foo", f("foo"), -1)
}

func Expand_test(){

	fmt.Println(Expand("xxxx$fooooxxxxxx$fooxx",strings.ToUpper))

}

func GetRootNode(url string) *html.Node{
	resp, err := http.Get(url)
	if err != nil{
		log.Fatalf("%v while accessing %v",err, url)
	}
	node, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err!= nil{
		log.Fatalf("%v while parsing response from %v",err,url)
	}
	return node



}

func ElementByID(n *html.Node, id string) *html.Node{
	if n.Type == html.ElementNode {
		for _, j := range n.Attr {
			if id == j.Key {
				fmt.Printf("%v---%v:\t%v\n",id,j.Key,j.Val)
				return n
			}
		}
	}
	for c:=n.FirstChild;c!=nil;c=c.NextSibling{
		  ElementByID(c,id)
	}
	return nil

}

func preChildren(n *html.Node, id string) bool{

	return ElementByID(n,id) != nil

}

//func postChildren(n *html.Node, id string)bool{
//
//	return false
//
//}

func ForEachNode(n *html.Node, pre, post func(*html.Node, string) bool, id string) *html.Node{

	if pre != nil && pre(n,id){
		return n
	}

	for c:=n.FirstChild;c!=nil;c=c.NextSibling{
		ForEachNode(c,pre,post,id)
	}

	if post != nil{
		post(n, id)
	}

	return nil

}


func main() {

	//Expand_test()
	rd := strings.NewReader(`<p><span><b width="sfe"">sfef</b><b width=265>2</b></span></p><b width="www"></b>`)
	node, _ := html.Parse(rd)
	n := ElementByID(node,"width")
        //node := ElementByID(GetRootNode("http://www.marathonguide.com"),"src")
	//node := ForEachNode(GetRootNode("http://www.marathonguide.com"), preChildren, nil, "src")
	if n != nil{
		fmt.Println(node.Data)
	}

}
