package main

import (
	"fmt"
	"os"
	"text/template"
)

type Todo struct {
	Name        string
	Description string
}

func main() {
	td := Todo{"Test templates", "Let's test a template to see the magic."}

	t, err := template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}
	fmt.Println("reuse:")
	tdNew := Todo{"Go", "Contribute to any Go project"}
	err = t.Execute(os.Stdout, tdNew)
}
