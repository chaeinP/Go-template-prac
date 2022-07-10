package main

import (
	"log"
	"os"
	"text/template"
)

func main(){
	/*
	PareFiles vs ParseGlob

	ParseFiles는 파일 단위로 파싱한다.
	ParseGlob은 폴더를 지정해 파싱할 수 있다.
	*/
	tpl, err:=template.ParseGlob("../../data/*.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)

	if err != nil {
		log.Fatalln(err)
	}
}
