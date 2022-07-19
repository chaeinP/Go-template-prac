package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type hotdog int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

/*
request body, url query data는 req.From에 담긴다.
request body는 req.PostForm에도 담긴다.
req.Form과 req.PostForm을 사용하기 위해선 반드시 req.ParseForm을 호출해야한다.
*/


func (m hotdog)ServeHTTP(w http.ResponseWriter, req *http.Request){
	err := req.ParseForm()
	if (err != nil){
		log.Fatalln(err)
	}
	fmt.Println(req.Form) // map[fname:[1234]]
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func main(){
	var d hotdog
	http.ListenAndServe(":8080", d);
}
