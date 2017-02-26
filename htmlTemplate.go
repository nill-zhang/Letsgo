package main

import (
	"net/http"
	"os"
	"fmt"
	"bufio"
	"encoding/json"
	"log"
	"html/template"
)

func GenerateTable(wr http.ResponseWriter, _ *http.Request){

	///////////////////////////////////////////////////////////////////////////////////////////////////
        template_str := `
        <html>
        <head>
        <style>
            table,th,td{border:0.2px solid black;
            border-collapse: collapse;}
        </style>
        </head>
        <body>
        <table style="width: 80%" align="center">

        <tr>
            <th>Company Name</th>
            <th>URL</th>
            <th>Official Twitter</th>
       	    <th>Category</th>
      	    <th>Description</th>
      	    <th>CreatedAt</th>
        </tr>
        {{range .}}
       	    <tr>
       	    <td>{{.Name}}</td>
       	    <td ><a href={{.Url}}>{{.Url}}</a></td>
       	    <td>{{.Twitter}}</td>
       	    <td>{{.Category}}</td>
       	    <td>{{.Description}}</td>
       	    <td>{{.CreatedAt}}</td>
        </tr>
        {{end}}
        </table>
        </body>
        `
        t := template.Must(template.New("company").Parse(template_str))
	type company struct {
		Name          string
		Url           string   `json:"homepage_url"`
		Twitter       string   `json:"twitter_username"`
		Category      string   `json:"category_code"`
		Description   string
		CreatedAt     string   `json:"created_at"`

	}
	f,err := os.Open("C:/Users/Admin/Documents/golang/resources/letsgo/companies.json")
	if err != nil{
		fmt.Fprintf(os.Stderr, "err %v", err)
	}
        s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	var cops []company

	var cop company

	for s.Scan(){


		err := json.Unmarshal(s.Bytes(),&cop)

		if err == nil{

			cops = append(cops, cop)
		}

	}

	t.Execute(wr, cops)

}


func StartHttpServer(){

	http.HandleFunc("/", GenerateTable)
	log.Fatal(http.ListenAndServe("localhost:800", nil))

}


func main() {

        StartHttpServer()

}
