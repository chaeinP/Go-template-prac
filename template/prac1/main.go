/*
	text-template 공식문서
	https://pkg.go.dev/text/template
*/

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	/*
	ParseFiles 원형
	func ParseFiles(filenames ...string)(*Template, error) {
		return parseFiles(nil, readFileOS, filenames...)
	}

	input으로 여러개의 string을 받을 수 있다.
	*/
	tpl, err := template.ParseFiles("../../data/tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	/*
	Execute 원형
	func (t *Template) Execute(wr io.Writer, data any) error {
		return t.execute(wr, data)
	}
	*/
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
