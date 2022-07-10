package main

import (
	"os"
	"text/template"
)

var tpl *template.Template
func init () {
	tpl =template.Must(template.ParseFiles("../../data/tpl1.gohtml"))
}

func main(){
	tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", "안녕")
}
