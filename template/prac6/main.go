package main

import (
	"os"
	"text/template"
)

var tpl *template.Template
func init () {
	tpl =template.Must(template.ParseFiles("../../data/tpl.gohtml"))
}

func main(){
	/*
	template에 동적으로 데이터를 주입시 데이터는 하나만 전달할 수 있다.

	따라서 map이나 struct형태로 데이터를 aggregate에 주입해아 한다.
	*/
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
}
