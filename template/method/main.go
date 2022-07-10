package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

/*
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Composition</title>
</head>
<body>

<h1>{{.Name}}</h1>
<h2>{{.Age}}</h2>
<h3>{{.SomeProcessing}}</h3> 7
<h3>{{.AgeDbl}}</h3> 112
<h3>{{.Age | .TakesArg}}</h3> 112
<h3>{{.AgeDbl | .TakesArg}}</h3> 224
</body>
</html>
*/

func main() {

	p := person{
		Name: "Ian Fleming",
		Age:  56,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}
