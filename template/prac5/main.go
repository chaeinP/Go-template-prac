package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

/*
	init 함수는 프로그램이 시작되면 한번만 실행된다.

	template.Must는 parseGlob의 반환 값 중 에러가 있는 지 검사하고 에러가 있으면 panic을 발생, err가 없으면 template만 반환한다.

	func Must(t *Template, err error) *Template {
		if err != nil {
			panic(err)
		}
	return t
	}
*/
func init(){
	tpl = template.Must(template.ParseGlob("../../data/*"))
}


func main(){
	tpl.Execute(os.Stdout, nil)
}
