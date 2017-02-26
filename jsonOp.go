package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"os"
	"log"
)

// OMG, you know what, if you want to use json to Marshal this struct's data
// you must make them exportable which means using Uppercase for the first letter of field name
type employee struct {
	Name       string   `json:"名字"` // json metadata tags no space after colon
	Age        uint8    `json:"年龄"`
	Weight     float32  `json:"体重"`
	Skills     []string `json:"特长"`
	Department string   `json:"部门"`
}

func UnmarshalJsonStrings() {

	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	// The basic JSON types are numbers (in decimal or scientific notation), boole ans (true or
	// false), and strings, which are sequences of Unicode co de points enclosed in double quotes,
	// with backslash escapes using a similar notation to Go, though JSON’s \Uhhhh numeric escapes
	// denote UTF-16 codes, not runes.
	// These basic typ es may be combined recursively using JSON arrays and objects
	// json key to struct field matching is case-insensitive, but
	// struct filed must be uppercase, for the Released index its a json tag in the movie struct
	// it will restored to the Year field
	data := []byte(`
    	{
        	"title": "Transformer",
      		"Released": 2015,
     		"Color": true,
    	    	"Actors":["Alex", "John", "Mary"],
   	     	"Review":{"Aisa":90,"America":86,"Europe":78}
    	}
	`)

	var m Movie
	err := json.Unmarshal(data, &m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v has occured")
	}
	fmt.Printf("%#v", m)

}

func MarshalUnmarshal() {

	var employees = []employee{
		{"sfzhang", 33, 70.5, []string{"c++", "mongodb", "redis", "golang"}, "embedded"},
		{"xlyang", 27, 71.5, []string{"python", "postgres", "redis", "golang"}, "devops"},
		{"qshu", 21, 75.5, []string{"excel", "django", "redis", "golang"}, "headquarters"},
		{"xlzhou", 42, 60.5, []string{"numpy", "pandas", "java", "golang"}, "logistics"},
		{"jchu", 27, 50.5, []string{"c++", "mysql", "perl", "ruby"}, "accounting"},
		{"lyzhang", 30, 10.5, []string{"php", "linux", "vmware", "php"}, "marketing"},
	}

	// Alternative:
	//var employees  = []employee{
	//	{Name:"sfzhang",Age:33, Weight:70.5,Skills:[]string{"c++", "mongodb", "redis", "golang"},Department:"embedded"},
	//        {Name:"xlyang",Age:27, Weight:71.5,Skills:[]string{"python", "postgres", "redis", "golang"},Department:"devops"},
	//        {Name:"qshu",Age:21, Weight:75.5,Skills:[]string{"excel", "django", "redis", "golang"},Department:"headquarters"},
	//	{Name:"xlzhou",Age:42, Weight:60.5,Skills:[]string{"numpy", "pandas", "java", "golang"},Department:"logistics"},
	//        {Name:"jchu",Age:27, Weight:50.5,Skills:[]string{"c++", "mysql", "perl", "ruby"},Department:"accounting"},
	//        {Name:"lyzhang",Age:30, Weight:10.5,Skills:[]string{"php", "linux", "vmware", "php"},Department:"marketing"},
	//        }

	// var movies = []Movie{
	// {Title: "Casablanca", Year: 1942, Color: false,
	//	 Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	// {Title: "Cool Hand Luke", Year: 1967, Color: true,
	//	 Actors: []string{"Paul Newman"}},
	// {Title: "Bullitt", Year: 1968, Color: true,
	//	 Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
	// }
	//res, err := json.Marshal(employees)
	// like python's pprint.pprint()
	res, err := json.MarshalIndent(employees, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", res)
	/////////////////////////////////////////////////////////////////////////////////////////////
	type emp1 struct {
		Name       string   `json:"名字"` // json metadata tags no space after colon
		Age        uint8    `json:"年龄"`
		Weight     float32  `json:"体重"`
		Skills     []string `json:"特长"`
		Department string   `json:"部门"`
	}

	var mp1 []emp1
	if json.Unmarshal(res, &mp1) == nil {
		fmt.Println(strings.Repeat("*", 80))
		fmt.Println("mp1: ", mp1)
	}
	////////////////////////////////////////////////////////////////////////////////////////////
	// partially unmarshal a json data
	// Any field in the Json data which doesn’t fit in the struct will just be ignored
	type emp2 struct {
		Name string `json:"名字"`
		Age  uint8  `json:"年龄"`
	}

	var mp2 []emp2
	if json.Unmarshal(res, &mp2) == nil {
		fmt.Println(strings.Repeat("*", 80))
		fmt.Println("mp2: ", mp2)
	}
	/////////////////////////////////////////////////////////////////////////////////////////////////

	var mp3 []struct {
		// the target struct field can be out of order compared to the original
		// struct
		Name       string   `json:"名字"`
		Department string   `json:"部门"`
		skills     []string `json:"特长"` // this will not be unmarshalled, because, its not exported
		Weight     float32  // this will also not be unmarshalled, because it missed
		// json tag `json:"体重"
		Age uint8 `json:"年龄"`
	}
	if json.Unmarshal(res, &mp3) == nil {
		fmt.Println(strings.Repeat("*", 80))
		fmt.Println("mp3: ", mp3)
	}
}

func DecodeEncode(jsonfile string){
	type company struct {
		Name string
		Url string `json:"homepage_url"`
		Tag string `json:"tag_list"`

	}
	var comp []company
	f, err := os.Open(jsonfile)
	if err != nil{
		log.Fatalf("%v while trying to open %v", err, jsonfile)

	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	if err = decoder.Decode(&comp); err != nil{
		log.Fatalf("%v while trying to decode", err)

	}

	for _,j := range comp{
		fmt.Printf("Name:%-20v \tUrl:%-30v \tTag:%-60v\n",j.Name, j.Url, j.Tag)
	}


}



func main() {

	//MarshalUnmarshal()
	//UnmarshalJsonStrings()
	DecodeEncode("C:/Users/Admin/Documents/golang/resources/letsgo/companies-2.json")

}
