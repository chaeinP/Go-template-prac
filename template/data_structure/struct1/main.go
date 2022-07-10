package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}


func init() {
	tpl = template.Must(template.ParseFiles("../../../data/struct1.gohtml"))
}

/*
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Peeps</title>
</head>
<body>
<ul>
    <li>{{.Name}}-{{.Motto}}</li>
</ul>
</body>
</html>
*/

func main() {

	buddha := sage{
		"Buddha",
		"The belief of no beliefs",
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
