package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml", "ywa.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	{
		tpl := template.Must(template.ParseFiles("ywa.html"))
		tpl.Execute(os.Stdout, nil)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
