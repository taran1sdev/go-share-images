package main

import (
	"html/template"
	"os"
)

// type User struct {
// 	Name string
// }

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := struct {
		Name string
		Bio  string // For escaping html //template.HTML For processing as html
	}{
		Name: "John Smith",
		Bio:  `<script>alert("xss")</script>`,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
