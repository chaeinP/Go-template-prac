package main

import (
	"log"
	"os"
	"text/template"
)

func main(){
	tpl, err := template.ParseFiles("../../data/two.gmao")

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("../../data/one.gmao","../../data/three.gmao" )
	if err != nil {
		log.Fatalln(err)
	}

	/*
	 Execute vs ExecuteTemplate

	 Execute는 tpl 이 담고 있는 텍스트중 첫번째만 실행한다.

	 ExecuteTemplate은 다수의 텍스트를 담고있는 tpl중 실행할 텍스트를 선택할 수 있다.
	*/
	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
